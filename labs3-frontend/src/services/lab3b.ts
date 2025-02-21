import { $apiStudentLab, $apiStudent } from 'src/boot/axios';
import { useServiceAction } from '../utils/service/action';
import { Lab } from 'src/models/lab/lab';
import { Lab3b } from 'src/models/lab/lab3b';

export const lab3bService = {
  getVariant: useServiceAction(() =>
    $apiStudentLab.get<Lab.GetVariant>(
      '/lab3b/variant?laboratory_id=1&minutes_duration=1000'
    )
  ),
  sendRuleNumber: useServiceAction((data: Lab3b.RuleNumber) =>
    $apiStudentLab.post<Lab3b.RuleNumber>('/lab3b/variant/rule-number', data)
  ),
  sendRuleValue: useServiceAction((data: unknown) =>
    $apiStudentLab.post('/lab3b/variant/rule-value', data)
  ),
  sendAllMatrices: useServiceAction((data: Lab3b.sendAllMatrices) =>
    $apiStudentLab.post<Lab3b.AllMatricesResult>(
      '/lab3b/variant/all-matrices',
      data
    )
  ),
  sendAllMatricesIntersection: useServiceAction(
    (data: Lab3b.sendAllMatricesIntersection) =>
      $apiStudentLab.post<Lab3b.sendAllMatricesIntersectionResult>(
        '/lab3b/variant/intersection',
        data
      )
  ),
  sendLevelSet: useServiceAction((data: Lab3b.SendLevelSet) =>
    $apiStudentLab.post<Lab3b.SendLevelSetResult>(
      '/lab3b/variant/level-set',
      data
    )
  ),
  sendCoff: useServiceAction((data: Lab.Coff) =>
    $apiStudentLab.post<Lab.CoffAnswer>('/lab3b/variant/coff-matrices', data)
  ),
  sendResult: useServiceAction((data: Lab.Result) =>
    $apiStudentLab.post<Lab.ResultAnswer>('/lab3b/variant/result', data)
  ),
  getInfo: useServiceAction(() =>
    $apiStudentLab.get<Lab.Info>('/lab3b/variant/info')
  ),
  increment0Step: useServiceAction(() =>
    $apiStudentLab.post<Lab.Info>('/lab3b/variant/increment-zero-step')
  ),
  increment2Step: useServiceAction(() =>
    $apiStudentLab.post<Lab.Info>('/lab3b/variant/increment-second-step')
  ),
};
