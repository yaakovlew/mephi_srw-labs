import { defineStore } from 'pinia';
import { Lab1a } from 'src/models/lab/lab1a';
import { lab1aService } from 'src/services/lab1a';
import { computed, ref } from 'vue';

export const useLab1aStore = defineStore('lab1a', () => {
  const info = ref<Lab1a.Info | null>(null);
  const variant = computed<Lab1a.Variant | null>(
    () => info.value?.variance.data ?? null
  );

  const saveVariant = async (newVariant: Lab1a.SaveVariant) => {
    if (variant.value) return;
    await lab1aService.saveVariant(newVariant);
    await updateInfo({ step: 1, percentage: 0 });
    await getInfo();
  };

  const getInfo = async () => {
    const res = await lab1aService.getInfo();
    if (res.data) {
      info.value = res.data;
    }
  };

  const updateInfo = async (data: Lab1a.UpdateInfo) => {
    await lab1aService.updateInfo(data);
  };

  const sendResult = async () => {
    await lab1aService.sendResult({ percentage: info.value?.percentage ?? 0 });
  };

  return {
    saveVariant,
    getInfo,
    variant,
    info,
    updateInfo,
    sendResult,
  };
});
