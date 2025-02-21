import { defineStore } from 'pinia';
import { Lab1a } from 'src/models/lab/lab1a';
import { Lab2 } from 'src/models/lab/lab2';
import { lab2Service } from 'src/services/lab2';
import { computed, ref } from 'vue';

export const useLab2Store = defineStore('lab2', () => {
  const info = ref<Lab2.Info | null>(null);
  const variant = computed<Lab2.Variant | null>(
    () => info.value?.variance.data ?? null
  );

  const saveVariant = async (newVariant: Lab2.SaveVariant) => {
    if (variant.value) return;
    await lab2Service.saveVariant(newVariant);
    await updateInfo({ step: 1, percentage: 0 });
    await getInfo();
  };

  const getInfo = async () => {
    const res = await lab2Service.getInfo();
    if (res.data) {
      info.value = res.data;
    }
  };

  const updateInfo = async (data: Lab1a.UpdateInfo) => {
    await lab2Service.updateInfo(data);
  };

  const sendResult = async () => {
    await lab2Service.sendResult({ percentage: info.value?.percentage ?? 0 });
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
