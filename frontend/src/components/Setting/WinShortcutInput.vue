<script setup lang="ts">
import { ref, computed, watch, onMounted, onBeforeUnmount } from 'vue';

const props = defineProps<{ modelValue: string }>();
const emit = defineEmits(['update:modelValue']);

const modifierKeys = ['Shift', 'Control', 'Alt', 'Meta'] as const;

const modifierLabels: Record<string, string> = {
  Shift: 'Shift',
  Control: 'Ctrl',
  Alt: 'Alt',
  Meta: 'Win', // Windows键
};

const normalizeModifier = (key: string): string => {
  switch (key.toUpperCase()) {
    case 'CMD':
    case 'COMMAND':
    case 'WIN':
    case 'WINDOWS':
      return 'Meta';
    case 'CTRL':
      return 'Control';
    case 'OPTION':
      return 'Alt';
    default:
      return key;
  }
};

const pressedModifiers = ref<Set<string>>(new Set());
const finalModifiers = ref<Set<string>>(new Set());
const mainKey = ref<string | null>(null);

const isRecording = ref(false);
const isFinalized = ref(false);

const display = computed(() => {
  const activeSet = isFinalized.value ? finalModifiers.value : pressedModifiers.value;

  const modifiers = modifierKeys
    .map(key => {
      const isActive = activeSet.has(key);
      const color = isActive ? '#1890ff' : '#999';
      return `<span style="color: ${color}; font-weight: bold;">${modifierLabels[key]}</span>`;
    })
    .join('  ');

  const keyColor = isFinalized.value && mainKey.value ? '#1890ff' : '#000';
  const keyPart = mainKey.value
    ? `<span style="margin-left: 6px; color: ${keyColor}; font-weight: bold;">${mainKey.value.toUpperCase()}</span>`
    : '';

  return modifiers + (mainKey.value ? '  ' : '') + keyPart;
});

const shortcutString = computed(() => {
  if (!isFinalized.value || !mainKey.value || finalModifiers.value.size === 0) {
    return '';
  }
  return [...finalModifiers.value].join('+') + '+' + mainKey.value.toUpperCase();
});

watch(shortcutString, (val) => {
  emit('update:modelValue', val);
});

const startRecording = () => {
  isRecording.value = true;
  isFinalized.value = false;
  mainKey.value = null;
  pressedModifiers.value.clear();
  finalModifiers.value.clear();
};

const finalize = () => {
  if (pressedModifiers.value.size === 0 || !mainKey.value) {
    startRecording();
    return;
  }
  isFinalized.value = true;
  finalModifiers.value = new Set(pressedModifiers.value);
};

const clearAll = () => {
  isRecording.value = false;
  isFinalized.value = false;
  pressedModifiers.value.clear();
  finalModifiers.value.clear();
  mainKey.value = null;
  emit('update:modelValue', '');
};

const onKeyDown = (e: KeyboardEvent) => {
  if (!isRecording.value || isFinalized.value) return;

  if (modifierKeys.includes(e.key as any)) {
    pressedModifiers.value.add(e.key);
  } else if (!['Tab', 'Escape'].includes(e.key)) {
    mainKey.value = e.key.length === 1 ? e.key.toUpperCase() : e.key;
    finalize();

    // 主键按下后立即取消焦点（触发 blur 逻辑）
    (document.activeElement as HTMLElement)?.blur();
  }
  e.preventDefault();
};

const onKeyUp = (e: KeyboardEvent) => {
  if (!isRecording.value || isFinalized.value) return;
  if (modifierKeys.includes(e.key as any)) {
    pressedModifiers.value.delete(e.key);
  }
};

onMounted(() => {
  window.addEventListener('keydown', onKeyDown);
  window.addEventListener('keyup', onKeyUp);
});
onBeforeUnmount(() => {
  window.removeEventListener('keydown', onKeyDown);
  window.removeEventListener('keyup', onKeyUp);
});

// 初始化值支持：从 modelValue 拆解 modifier 和主键
watch(() => props.modelValue, (val) => {
  if (!val) {
    clearAll();
    return;
  }

  const parts = val.split('+').map(normalizeModifier);
  const mods = parts.filter(p => modifierKeys.includes(p as any));
  const key = parts.find(p => !modifierKeys.includes(p as any));

  if (mods.length && key) {
    finalModifiers.value = new Set(mods);
    mainKey.value = key;
    isFinalized.value = true;
  } else {
    clearAll(); // 防止非法值
  }
}, { immediate: true });
</script>

<template>
  <a-input
    style="width: 160px;"
    readonly
    :value="''"
    @focus="startRecording"
    @mousedown.stop
    @blur="isRecording = false"
  >
    <template #suffix>
      <span v-html="display" />
    </template>
  </a-input>
</template>