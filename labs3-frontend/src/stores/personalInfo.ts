import { defineStore } from 'pinia';
import { IPersonalInfo } from 'src/models/personal-info/personalInfo';
import { PersonalInfoService } from 'src/services/personalInfo';
import { Ref, ref } from 'vue';

export const usePersonalInfoStore = defineStore('personal-info', () => {
  const me: Ref<IPersonalInfo | null> = ref(null);

  const getPersonalInfo = async () => {
    const res = await PersonalInfoService.fetch();
    if (res.data) {
      me.value = res.data;
    }
  };

  const changeName = async (name: string) => {
    await PersonalInfoService.changeName({ name });
  };

  const changeSurname = async (surname: string) => {
    await PersonalInfoService.changeSurname({ surname });
  };

  const changePassword = async (old_password: string, new_password: string) => {
    await PersonalInfoService.changePassword({ old_password, new_password });
  };

  return {
    me,
    getPersonalInfo,
    changeName,
    changeSurname,
    changePassword,
  };
});
