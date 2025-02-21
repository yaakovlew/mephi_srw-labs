import { $apiLecturer } from 'src/boot/axios';
import { IStudent } from 'src/models/student/student';
import { useServiceAction } from 'src/utils/service/action';

export const StudentService = {
  fetch: useServiceAction(() =>
    $apiLecturer.get<IStudent.FetchData>('/student')
  ),
  changeGroup: useServiceAction((data: IStudent.ChangeGroup) =>
    $apiLecturer.put<IStudent[]>('/student', data)
  ),
  changePassword: useServiceAction((data: IStudent.ChangePassword) =>
    $apiLecturer.put<IStudent[]>('/student/change-password', data)
  ),
};
