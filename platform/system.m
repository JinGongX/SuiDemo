#import <Cocoa/Cocoa.h>
#import <Carbon/Carbon.h>
#import <stdlib.h>
#import <AppKit/AppKit.h>
#import <ApplicationServices/ApplicationServices.h>
#import <Vision/Vision.h>
#import <Foundation/Foundation.h>
#import "system.h"

// ============ 热键相关 ============
// 声明 Go 回调函数
extern void hotKeyCallback(unsigned int keyCode, unsigned int modifiers);
// 定义热键最大数量
#define MAX_HOTKEYS 10
// 保存注册的 HotKeyRef
static EventHotKeyRef hotKeyRefs[MAX_HOTKEYS];
static EventHotKeyID hotKeyIDs[MAX_HOTKEYS];
static int hotKeyCount = 0;
static int handlerInstalled = 0;

OSStatus MyHotKeyHandler(EventHandlerCallRef nextHandler, EventRef theEvent, void *userData) {
    EventHotKeyID hkID;
    GetEventParameter(theEvent, kEventParamDirectObject, typeEventHotKeyID, NULL, sizeof(hkID), NULL, &hkID);

    unsigned int keyCode = hkID.id >> 16;
    unsigned int modifiers = hkID.id & 0xFFFF;

    hotKeyCallback(keyCode, modifiers); // ✅ 正确传值给 Go
    return noErr;
}

void InstallHotKeyHandler() {
    if (handlerInstalled) return;

    EventTypeSpec eventType;
    eventType.eventClass = kEventClassKeyboard;
    eventType.eventKind = kEventHotKeyPressed;

    InstallApplicationEventHandler(&MyHotKeyHandler, 1, &eventType, NULL, NULL);
    handlerInstalled = 1;
}

void RegisterHotKeyDynamic(unsigned int keyCode, unsigned int modifiers) {
    InstallHotKeyHandler();

    if (hotKeyCount >= MAX_HOTKEYS) return;

    EventHotKeyID hkID;
    hkID.signature = 'htk1';
    hkID.id = (keyCode << 16) | (modifiers & 0xFFFF); // ✅ 编码 keyCode 和 modifiers

    RegisterEventHotKey(keyCode, modifiers, hkID, GetApplicationEventTarget(), 0, &hotKeyRefs[hotKeyCount]);
    // ✅ 保存 ID
    hotKeyIDs[hotKeyCount] = hkID;
    hotKeyCount++;
}
void NSAppActivateIgnoringOtherApps() {
    [[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateIgnoringOtherApps];
}

void UnregisterHotKey(unsigned int keyCode, unsigned int modifiers) {
    unsigned int targetID = (keyCode << 16) | (modifiers & 0xFFFF);

    for (int i = 0; i < hotKeyCount; i++) {
        unsigned int existingID = hotKeyIDs[i].id;
        if (existingID == targetID) {
            // 注销
            UnregisterEventHotKey(hotKeyRefs[i]);

            // 清理数组（用后面的项覆盖）
            for (int j = i; j < hotKeyCount - 1; j++) {
                hotKeyRefs[j] = hotKeyRefs[j + 1];
                hotKeyIDs[j] = hotKeyIDs[j + 1];
            }
            hotKeyCount--;
            return;
        }
    }
}

void HideDockIcon(void) {
    NSApplication *app = [NSApplication sharedApplication];
    [app setActivationPolicy:NSApplicationActivationPolicyProhibited];
}


// ============ OCR功能 ============
const char* VisionOCR(const char* path) {
    @autoreleasepool {
        NSString* imagePath = [NSString stringWithUTF8String:path];
        NSData* imageData = [NSData dataWithContentsOfFile:imagePath];
        NSImage* image = [[NSImage alloc] initWithData:imageData];
        if (!image) return "Image load failed";

        CGImageRef cgRef = [image CGImageForProposedRect:NULL context:nil hints:nil];
        if (!cgRef) return "CGImage conversion failed";

        VNRecognizeTextRequest *request = [[VNRecognizeTextRequest alloc] init];
        request.recognitionLevel = VNRequestTextRecognitionLevelAccurate;

        VNImageRequestHandler *handler = [[VNImageRequestHandler alloc] initWithCGImage:cgRef options:@{}];
        NSError *error = nil;
        [handler performRequests:@[request] error:&error];

        if (error) return [[error.localizedDescription UTF8String] copy];

        NSMutableString *result = [NSMutableString string];
        for (VNRecognizedTextObservation *obs in request.results) {
            VNRecognizedText *topCandidate = [[obs topCandidates:1] firstObject];
            if (topCandidate) {
                [result appendString:topCandidate.string];
                [result appendString:@"\n"];
            }
        }

        return strdup([result UTF8String]);
    }
}

const char *VisionOCRFromMemory(const void *data, int size) {
    @autoreleasepool {
        NSData *imageData = [NSData dataWithBytes:data length:size];
        NSImage *image = [[NSImage alloc] initWithData:imageData];
        CGImageRef cgRef = [image CGImageForProposedRect:nil context:nil hints:nil];
        VNRecognizeTextRequest *req = [[VNRecognizeTextRequest alloc] init];
        req.recognitionLevel = VNRequestTextRecognitionLevelAccurate;
        req.recognitionLanguages = @[@"zh-Hans", @"en-US"];
        VNImageRequestHandler *handler = [[VNImageRequestHandler alloc] initWithCGImage:cgRef options:@{}];
        NSError *err = nil;
        [handler performRequests:@[req] error:&err];
        if (err) return strdup([[err localizedDescription] UTF8String]);
        NSMutableString *out = [NSMutableString string];
        for (VNRecognizedTextObservation *obs in req.results) {
            VNRecognizedText *cand = [[obs topCandidates:1] firstObject];
            if (cand) {
                [out appendString:cand.string];
                [out appendString:@"\n"];
            }
        }
        return strdup([out UTF8String]);
    }
}