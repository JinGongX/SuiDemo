// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import { Call as $Call, CancellablePromise as $CancellablePromise, Create as $Create } from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as application$0 from "../github.com/wailsapp/wails/v3/pkg/application/models.js";

export function OpenSecondWindow(): $CancellablePromise<void> {
    return $Call.ByID(1306260514);
}

export function SetApp(app: application$0.App | null): $CancellablePromise<void> {
    return $Call.ByID(3487267257, app);
}
