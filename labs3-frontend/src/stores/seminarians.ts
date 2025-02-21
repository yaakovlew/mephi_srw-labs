import { defineStore } from 'pinia';
import { ISeminarian } from 'src/models/seminarian/seminarian';
import { SeminariansService } from 'src/services/seminarians';
import { Ref, ref } from 'vue';

export const useSeminariansStore = defineStore('seminarians', () => {
  const seminarians: Ref<ISeminarian.Seminarian[] | null> = ref(null);

  const fetchSeminarians = async () => {
    const res = await SeminariansService.fetch();
    console.log(res);
    if (res.data) {
      seminarians.value = res.data.seminarians;
    }
  };

  return {
    seminarians,
    fetchSeminarians,
  };
});
