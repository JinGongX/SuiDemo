import {GetOS} from "../../bindings/changeme/internal/services/systeminfo";

let cachedOS: string | null = null;
let osInitPromise: Promise<void>
/**
 * 初始化 OS，只调用一次
 */
export  function initOS(): Promise<void> {
 if (!osInitPromise) {
    osInitPromise = GetOS().then(os => {
      cachedOS = os
    })
  }
  return osInitPromise;
}


export function getOS(): string | null {
  if (!cachedOS) {
    throw new Error("OS not initialized yet, await OS_READY first")
  }
  return cachedOS;
}

export function IsmacOS(): boolean {
  return cachedOS==='darwin';
}

// ✅ 导出一个初始化状态 Promise
export const OS_READY = initOS()