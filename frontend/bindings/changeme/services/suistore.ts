// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import { Call as $Call, CancellablePromise as $CancellablePromise, Create as $Create } from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function Close(): $CancellablePromise<void> {
    return $Call.ByID(1355592200);
}

export function GetHotkeys(): $CancellablePromise<$models.Hotkey[]> {
    return $Call.ByID(23462219).then(($result: any) => {
        return $$createType1($result);
    });
}

export function Start(): $CancellablePromise<void> {
    return $Call.ByID(835708248);
}

export function Stop(): $CancellablePromise<void> {
    return $Call.ByID(1501773308);
}

/**
 * 快捷键修改
 */
export function UpHotkey(id: number, key: number, modifier: number): $CancellablePromise<void> {
    return $Call.ByID(3795952175, id, key, modifier);
}

// Private type creation functions
const $$createType0 = $models.Hotkey.createFrom;
const $$createType1 = $Create.Array($$createType0);
