import { $apiLecturer, $apiSemianrian, $apiStudent } from 'src/boot/axios';
import { IDiscipline } from 'src/models/discipline/discipline';
import { IGroup } from 'src/models/group/group';
import { useAuthStore } from 'src/stores/auth';
import { useServiceAction } from 'src/utils/service/action';
import { computed } from 'vue';

export const DisciplineService = {
  fetch: useServiceAction(() => {
    const authStore = useAuthStore();
    const userType = computed(() => authStore.userType);
    switch (userType.value) {
      case 'student':
        return $apiStudent.get<IDiscipline.DisciplineFetchData>('/disciplines');
      case 'lecturer':
        return $apiLecturer.get<IDiscipline.DisciplineFetchData>('/discipline');
      case 'seminarian':
        return $apiSemianrian.get<IDiscipline.DisciplineFetchData>(
          '/discipline'
        );
      default:
        break;
    }
    return $apiStudent.get<IDiscipline.DisciplineFetchData>('/disciplines');
  }),
  getAllGroups: useServiceAction((data: IDiscipline.FetchData) =>
    $apiLecturer.get<IGroup.GroupsFetchData>(`/discipline/group/${data.id}`)
  ),
  getAlvaliableGroups: useServiceAction((data: IDiscipline.FetchData) =>
    $apiLecturer.get<IGroup.GroupsFetchData>(
      `/discipline/group/available-to-add/${data.id}`
    )
  ),
  changeGroupName: useServiceAction((data: IDiscipline.Discipline) =>
    $apiLecturer.put<IGroup.GroupsFetchData>('/discipline', data)
  ),
  addGroupToDiscipline: useServiceAction((data: IDiscipline.AddGroup) =>
    $apiLecturer.post<IDiscipline.DisciplineFetchData>(
      '/discipline/group',
      data
    )
  ),
  deleteGroupFromDiscipline: useServiceAction((data: IDiscipline.AddGroup) =>
    $apiLecturer.delete<IDiscipline.DisciplineFetchData>('/discipline/group', {
      data,
    })
  ),
  getDisciplineInfo: useServiceAction((data: IDiscipline.FetchData) =>
    $apiLecturer.get<IDiscipline.DisciplineFullInfo>(`/discipline/${data.id}`)
  ),
  createDiscipline: useServiceAction((data: IDiscipline.CreateDiscipline) =>
    $apiLecturer.post('/discipline', data)
  ),
  changeDisciplineName: useServiceAction(
    (data: IDiscipline.ChangeDisciplineName) =>
      $apiLecturer.put('/discipline', data)
  ),
  changeDisciplineExamMark: useServiceAction(
    (data: IDiscipline.ChangeDisciplineExamMark) =>
      $apiLecturer.put('/discipline/mark/exam', data)
  ),
  changeDisciplineLessonMark: useServiceAction(
    (data: IDiscipline.ChangeDisciplineLessonMark) =>
      $apiLecturer.put('/discipline/mark/lesson', data)
  ),
  changeDisciplineSeminarMark: useServiceAction(
    (data: IDiscipline.ChangeDisciplineSeminarMark) =>
      $apiLecturer.put('/discipline/mark/seminar', data)
  ),
  getGroupSeminarians: useServiceAction((body: IGroup.FetchGroupSeminarians) =>
    $apiLecturer.get<IGroup.GroupSeminarians>(
      `/group/seminarian?discipline_id=${body.discipline_id}&group_id=${body.group_id}`
    )
  ),
  addSeminarianToGroup: useServiceAction((data: IGroup.AddSeminarian) => {
    return $apiLecturer.post<IGroup.GroupSeminarians>('/seminarian', data);
  }),
  // deleteSeminarianFromGroup: useServiceAction((data: IGroup.AddSeminarian) => {
  //   return $apiLecturer.delete<IGroup.GroupSeminarians>(
  //     '/group/seminarian',
  //     data
  //   );
  // }),
};
