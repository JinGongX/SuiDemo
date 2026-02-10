// hotkeyUtils.ts

const modifierMap: Record<string, number> = {
  Control: 256,
  Alt: 512,
  Shift: 1024,
  Meta: 2048,
};

// 修饰键位定义（macOS 通常使用）
export const MODIFIERS = {
  CMD: 1 << 8,     // 256
  SHIFT: 1 << 9,   // 512
  ALT: 1 << 10,    // 1024
  CTRL: 1 << 11,   // 2048
};

// 快捷键字符串 => 键码映射（mac keycode 映射）
const keyMap: Record<string, number> = {
  A: 0, B: 11, C: 8, D: 2, E: 14, F: 3, G: 5,
  H: 4, I: 34, J: 38, K: 40, L: 37, M: 46,
  N: 45, O: 31, P: 35, Q: 12, R: 15, S: 1,
  T: 17, U: 32, V: 9, W: 13, X: 7, Y: 16, Z: 6,
};


export function parseHotkeyString(hotkey: string): { key: number, modifier: number } | null {
  if (!hotkey.includes('+')) return null;

  const parts = hotkey.split('+');
  const mods = parts.slice(0, -1); // 修饰键
  const keyStr = parts[parts.length - 1];

  let modifier = 0;
  for (const mod of mods) {
    if (modifierMap[mod]) modifier |= modifierMap[mod];
  }

  const key = keyStr.toUpperCase().charCodeAt(0); // A → 65, B → 66, S → 83 ...
  if (!key || isNaN(key)) return null;

  return { key, modifier };
}

export function parseShortcutToHotkey(shortcut: string): { key: number, modifier: number } {
  const parts = shortcut.toUpperCase().split('+');
  let keys = '';
  let modifier = 0;

  for (const part of parts) {
    switch (part) {
      case 'CMD':
      case 'COMMAND':
      case 'META':
        modifier |= 1 << 8;
        break;
      case 'SHIFT':
        modifier |= 1 << 9;
        break;
      case 'ALT':
      case 'OPTION':
        modifier |= 1 << 10;
        break;
      case 'CTRL':
      case 'CONTROL':
        modifier |= 1 << 11;
        break;
      default:
        keys = part;
    }
  }

  // const keyMap: Record<string, number> = {
  //   A: 0, B: 11, C: 8, D: 2, E: 14, F: 3, G: 5,
  //   H: 4, I: 34, J: 38, K: 40, L: 37, M: 46,
  //   N: 45, O: 31, P: 35, Q: 12, R: 15, S: 1,
  //   T: 17, U: 32, V: 9, W: 13, X: 7, Y: 16, Z: 6,
  // };

  const key = keyMap[keys];
  return {
    key,
    modifier,
  };
}

// 反向映射：keycode => 字符
const reverseKeyMap = Object.fromEntries(
  Object.entries(keyMap).map(([k, v]) => [v, k])
);
// 将 keycode 和 modifier 转换为字符串（如 40, 768 => "Cmd+Shift+K"）
export function formatHotkeyString(key: number, modifier: number): string {
  const parts: string[] = [];

  if (modifier & MODIFIERS.CMD) parts.push('Cmd');
  if (modifier & MODIFIERS.SHIFT) parts.push('Shift');
  if (modifier & MODIFIERS.ALT) parts.push('Alt');
  if (modifier & MODIFIERS.CTRL) parts.push('Ctrl');

  const keyStr = reverseKeyMap[key];
  if (!keyStr) throw new Error(`Unknown keycode: ${key}`);

  parts.push(keyStr.toUpperCase());
  return parts.join('+');
}

export function formatHotkeyStringmac(keycode: number, modifiers: number): string {
  const keyMap: Record<number, string> = {
    0: 'A', 11: 'B', 8: 'C', 2: 'D', 14: 'E', 3: 'F', 5: 'G',
    4: 'H', 34: 'I', 38: 'J', 40: 'K', 37: 'L', 46: 'M',
    45: 'N', 31: 'O', 35: 'P', 12: 'Q', 15: 'R', 1: 'S',
    17: 'T', 32: 'U', 9: 'V', 13: 'W', 7: 'X', 16: 'Y', 6: 'Z'
  };

  const mods: string[] = [];

  if (modifiers & (1 << 8)) mods.push('Cmd');      // mac
  if (modifiers & (1 << 9)) mods.push('Shift');
  if (modifiers & (1 << 10)) mods.push('Alt');
  if (modifiers & (1 << 11)) mods.push('Control');

  const key = keyMap[keycode] ?? 'Unknown';

  return [...mods, key].join('+');
}

//win
export function formatHotkeyStringWin(keycode: number, modifiers: number): string {
  const keyMap: Record<number, string> = {
    65: 'A', 66: 'B', 67: 'C', 68: 'D', 69: 'E', 70: 'F', 71: 'G',
    72: 'H', 73: 'I', 74: 'J', 75: 'K', 76: 'L', 77: 'M', 78: 'N',
    79: 'O', 80: 'P', 81: 'Q', 82: 'R', 83: 'S', 84: 'T', 85: 'U',
    86: 'V', 87: 'W', 88: 'X', 89: 'Y', 90: 'Z',
    13: 'Enter', 27: 'Esc', 32: 'Space',
    112: 'F1', 113: 'F2', 114: 'F3', 115: 'F4', 116: 'F5', 117: 'F6',
    118: 'F7', 119: 'F8', 120: 'F9', 121: 'F10', 122: 'F11', 123: 'F12',
    37: 'Left', 38: 'Up', 39: 'Right', 40: 'Down',
  };

  const mods: string[] = [];

  if (modifiers & 1) mods.push('Alt');
  if (modifiers & 2) mods.push('Ctrl');
  if (modifiers & 4) mods.push('Shift');
  if (modifiers & 8) mods.push('Win');

  const key = keyMap[keycode] ?? `KeyCode(${keycode})`;

  return [...mods, key].join('+');
}


export function parseShortcutToHotkeyWin(shortcut: string): { key: number, modifier: number } {
  const parts = shortcut.toUpperCase().split('+');
  let keyPart = '';
  let modifier = 0;

  for (const part of parts) {
    switch (part) {
      case 'ALT':
        modifier |= 1;
        break;
      case 'CTRL':
      case 'CONTROL':
        modifier |= 2;
        break;
      case 'SHIFT':
        modifier |= 4;
        break;
      case 'WIN':
      case 'CMD':
      case 'META':
        modifier |= 8;
        break;
      default:
        keyPart = part;
    }
  }
 
  const key = keyWinMap[keyPart] ?? parseInt(keyPart);

  return {
    key,
    modifier,
  };
}

// 快捷键字符串 => 键码映射（win keycode 映射）
const keyWinMap: Record<string, number> = {
  A: 65, B: 66, C: 67, D: 68, E: 69, F: 70, G: 71,
    H: 72, I: 73, J: 74, K: 75, L: 76, M: 77, N: 78,
    O: 79, P: 80, Q: 81, R: 82, S: 83, T: 84, U: 85,
    V: 86, W: 87, X: 88, Y: 89, Z: 90,
    ENTER: 13, ESC: 27, SPACE: 32,
    F1: 112, F2: 113, F3: 114, F4: 115,
    F5: 116, F6: 117, F7: 118, F8: 119,
    F9: 120, F10: 121, F11: 122, F12: 123,
    UP: 38, DOWN: 40, LEFT: 37, RIGHT: 39,
};