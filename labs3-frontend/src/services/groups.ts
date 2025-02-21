import { $apiLecturer, $apiPublic } from 'src/boot/axios';
import { IDiscipline } from 'src/models/discipline/discipline';
import { IGroup } from 'src/models/group/group';
import { ISeminarian } from 'src/models/seminarian/seminarian';
import { useServiceAction } from 'src/utils/service/action';

export const GroupsService = {
  fetch: useServiceAction(() =>
    $apiLecturer.get<IGroup.GroupsFetchData>('/group')
  ),
  fetchCommon: useServiceAction(() =>
    $apiPublic.get<IGroup.GroupsFetchData>('/groups')
  ),
  addNewGroup: useServiceAction((data: IGroup.AddGroup) =>
    $apiLecturer.post<IGroup.GroupsFetchData>('/group', data)
  ),
  changeGroupName: useServiceAction((data: IGroup.ChangeGroupName) =>
    $apiLecturer.put<IGroup.GroupsFetchData>('/group', data)
  ),
  deleteGroup: useServiceAction((data: IGroup.DeleteGroup) =>
    $apiLecturer.delete<IGroup.GroupsFetchData>(`/group/${data.id}`)
  ),
  getGroupStudents: useServiceAction((data: IGroup.FetchGroupStudents) =>
    $apiLecturer.get<IGroup.GetGroupStudentStudentId>(
      `/group/students/${data.id}`
    )
  ),
  getGroupDiscipline: useServiceAction((data: IGroup.FetchGroupStudents) =>
    $apiLecturer.get<IDiscipline.DisciplineFetchData>(
      `/group/discipline/${data.id}`
    )
  ),
  getGroupSeminarians: useServiceAction((body: IGroup.FetchGroupSeminarians) =>
    $apiLecturer.get<IGroup.GroupSeminarians>(
      `/group/seminarian?discipline_id=${body.discipline_id}&group_id=${body.group_id}`
    )
  ),
};
