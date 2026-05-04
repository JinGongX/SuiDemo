#ifndef CLIPBOARD_H
#define CLIPBOARD_H
#include <stdbool.h>
 
// 热键相关函数
void RegisterHotKeyDynamic(unsigned int keyCode, unsigned int modifiers);
void InstallHotKeyHandler(void);
void NSAppActivateIgnoringOtherApps();  // ✅ 只声明，不要实现！
void UnregisterHotKey(unsigned int keyCode, unsigned int modifiers);

void HideDockIcon(void);
// OCR功能
const char* VisionOCR(const char* path);//ocr
const char *VisionOCRFromMemory(const void *data, int size);
#endif