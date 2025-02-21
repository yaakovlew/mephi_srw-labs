import { $apiStudentLab12 } from 'src/boot/axios';
import { useServiceAction } from '../utils/service/action';
import { Lab1a } from 'src/models/lab/lab1a';
import { Lab2 } from 'src/models/lab/lab2';

export const lab2Service = {
  saveVariant: useServiceAction((data: Lab2.SaveVariant) =>
    $apiStudentLab12.post<Lab1a.SaveVariant>('/lab2/variant', data)
  ),
  getInfo: useServiceAction(() =>
    $apiStudentLab12.get<Lab2.Info>('/lab2/variant/info')
  ),
  updateInfo: useServiceAction((data: Lab1a.UpdateInfo) =>
    $apiStudentLab12.post<Lab1a.Info>('/lab2/variant/info', data)
  ),
  sendResult: useServiceAction((data: Lab1a.Result) =>
    $apiStudentLab12.post<Lab1a.Result>('/lab2/variant/result', data)
  ),
};
