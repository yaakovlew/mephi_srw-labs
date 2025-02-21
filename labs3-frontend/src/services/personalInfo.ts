import { $apiLecturer, $apiSemianrian, $apiStudent } from 'src/boot/axios';
import { IPersonalInfo } from 'src/models/personal-info/personalInfo';
import { useAuthStore } from 'src/stores/auth';
import { useServiceAction } from 'src/utils/service/action';
import { computed } from 'vue';

export const PersonalInfoService = {
  fetch: useServiceAction(() => {
    const authStore = useAuthStore();
    const userType = computed(() => authStore.userType);

    switch (userType.value) {
      case 'student':
        return $apiStudent.get<IPersonalInfo>('/personal-data');
      case 'lecturer':
        return $apiLecturer.get<IPersonalInfo>('/personal-data');
      case 'seminarian':
        console.log('Should be here', $apiSemianrian);
        return $apiSemianrian.get<IPersonalInfo>('/personal-data');
      default:
        break;
    }
    return $apiStudent.get<IPersonalInfo>('/personal-data');
  }),

  changeName: useServiceAction((data: IPersonalInfo.IChangeName) => {
    const authStore = useAuthStore();
    const userType = computed(() => authStore.userType);
    console.log(userType.value);
    switch (userType.value) {
      case 'student':
        return $apiStudent.put<IPersonalInfo>('/personal-data/name', data);
      case 'lecturer':
        return $apiLecturer.put<IPersonalInfo>('/personal-data/name', data);
      case 'seminarian':
        return $apiSemianrian.put<IPersonalInfo>('/personal-data/name', data);
      default:
        break;
    }
    return $apiStudent.put<IPersonalInfo>('/personal-data/name', data);
  }),

  changeSurname: useServiceAction((data: IPersonalInfo.IChangeSurname) => {
    const authStore = useAuthStore();
    const userType = computed(() => authStore.userType);
    switch (userType.value) {
      case 'student':
        return $apiStudent.put<IPersonalInfo>('/personal-data/surname', data);
      case 'lecturer':
        return $apiLecturer.put<IPersonalInfo>('/personal-data/surname', data);
      case 'seminarian':
        return $apiSemianrian.put<IPersonalInfo>(
          '/personal-data/surname',
          data
        );
      default:
        break;
    }
    return $apiStudent.put<IPersonalInfo>('/personal-data/surname', data);
  }),

  changePassword: useServiceAction((data: IPersonalInfo.IChangePassword) => {
    const authStore = useAuthStore();
    const userType = computed(() => authStore.userType);
    switch (userType.value) {
      case 'student':
        return $apiStudent.put<IPersonalInfo>(
          '/personal-data/change-password',
          data
        );
      case 'lecturer':
        return $apiLecturer.put<IPersonalInfo>(
          '/personal-data/change-password',
          data
        );
      case 'seminarian':
        return $apiSemianrian.put<IPersonalInfo>(
          '/personal-data/change-password',
          data
        );
      default:
        break;
    }
    return $apiStudent.put<IPersonalInfo>(
      '/personal-data/change-password',
      data
    );
  }),
};
