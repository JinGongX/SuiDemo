import {GetOS} from "../../bindings/changeme/internal/services/systeminfo";

let cachedOS: string | null = null;
let osInitPromise: Promise<void> | null = null
/**
 * 初始化 OS，只调用一次
 */
export async function initOS() {
//   if (!cachedOS) {
//     cachedOS = await GetOS();
//   }
 if (!osInitPromise) {
    osInitPromise = GetOS().then(os => {
      cachedOS = os
    })
  }
  return osInitPromise
}

export function getOS(): string | null {
  return cachedOS;
}

export function IsmacOS(): boolean {
  return cachedOS==='darwin';
}

// ✅ 导出一个初始化状态 Promise
export const OS_READY: Promise<void> = initOS()