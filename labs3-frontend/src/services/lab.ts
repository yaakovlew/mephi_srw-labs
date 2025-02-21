import { $apiStudentLab } from 'src/boot/axios';
import { useServiceAction } from '../utils/service/action';
import { Lab } from 'src/models/lab/lab';

export const labService = {
  getVariant: useServiceAction(() =>
    $apiStudentLab.get<Lab.GetVariant>(
      '/lab3a/variant?laboratory_id=1&minutes_duration=1000'
    )
  ),
  sendAlternativeSet: useServiceAction((data: Lab.AlternativeSet) =>
    $apiStudentLab.post<Lab.AlternativeSetAnswer>(
      '/lab3a/variant/alternative-sets',
      data
    )
  ),
  sendSecondNonDominated: useServiceAction((data: Lab.NonDominated) =>
    $apiStudentLab.post<Lab.NonDominatedAnswer>(
      '/lab3a/variant/second-non-dominated',
      data
    )
  ),
  sendNonDominated: useServiceAction((data: Lab.NonDominated) =>
    $apiStudentLab.post<Lab.NonDominatedAnswer>(
      '/lab3a/variant/non-dominated',
      data
    )
  ),
  sendAlternativeDiffMatrices: useServiceAction(
    (data: Lab.AlternativeDiffMatrices) =>
      $apiStudentLab.post<Lab.AlternativeDiffMatricesAnswer>(
        '/lab3a/variant/diff-matrices',
        data
      )
  ),
  sendIntersection: useServiceAction((data: Lab.Intersection) =>
    $apiStudentLab.post<Lab.IntersectionAnswer>(
      '/lab3a/variant/intersection',
      data
    )
  ),
  sendCoff: useServiceAction((data: Lab.Coff) =>
    $apiStudentLab.post<Lab.CoffAnswer>('/lab3a/variant/coff-matrices', data)
  ),
  sendResult: useServiceAction((data: Lab.Result) =>
    $apiStudentLab.post<Lab.ResultAnswer>('/lab3a/variant/result', data)
  ),
  getInfo: useServiceAction(() =>
    $apiStudentLab.get<Lab.Info>('/lab3a/variant/info')
  ),
  incrementSecondStep: useServiceAction(() =>
    $apiStudentLab.post<Lab.Info>('/lab3a/variant/increment-second-step')
  ),
};
