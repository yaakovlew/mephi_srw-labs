import { $apiStudentLab12 } from 'src/boot/axios';
import { useServiceAction } from '../utils/service/action';
import { Lab1a } from 'src/models/lab/lab1a';

export const lab1aService = {
  saveVariant: useServiceAction((data: Lab1a.SaveVariant) =>
    $apiStudentLab12.post<Lab1a.SaveVariant>('/lab1a/variant', data)
  ),
  getInfo: useServiceAction(() =>
    $apiStudentLab12.get<Lab1a.Info>('/lab1a/variant/info')
  ),
  updateInfo: useServiceAction((data: Lab1a.UpdateInfo) =>
    $apiStudentLab12.post<Lab1a.Info>('/lab1a/variant/info', data)
  ),
  sendResult: useServiceAction((data: Lab1a.Result) =>
    $apiStudentLab12.post<Lab1a.Result>('/lab1a/variant/result', data)
  ),
};
