import { $apiStudentLab, $apiStudent } from 'src/boot/axios';
import { useServiceAction } from '../utils/service/action';
import { Lab } from 'src/models/lab/lab';
import { Lab3c } from 'src/models/lab/lab3c';

export const lab3cService = {
  getVariant: useServiceAction(() =>
    $apiStudentLab.get<Lab.GetVariant>(
      '/lab3c/variant?laboratory_id=1&minutes_duration=1000'
    )
  ),
  sendAlternativeMatrix: useServiceAction((data: Lab3c.AlternativeMatrix) =>
    $apiStudentLab.post<Lab3c.AlternativeMatrixResult>(
      '/lab3c/variant/alternative-matrix',
      data
    )
  ),
  sendCriteriaMatrix: useServiceAction((data: Lab3c.CriteriaMatirx) =>
    $apiStudentLab.post<Lab3c.CriteriaMatirxResult>(
      '/lab3c/variant/criteria-matrix',
      data
    )
  ),
  sendAlternativeMatrices: useServiceAction((data: Lab3c.AlternativeMatrices) =>
    $apiStudentLab.post<Lab3c.AlternativeMatricesResult>(
      '/lab3c/variant/current-matrix',
      data
    )
  ),
  sendCriteriaEstimation: useServiceAction((data: Lab3c.CriteriaEstimation) =>
    $apiStudentLab.post<Lab3c.CriteriaEstimationResult>(
      '/lab3c/variant/estimation',
      data
    )
  ),
  sendEstimation: useServiceAction((data: Lab3c.CriteriaEstimation) =>
    $apiStudentLab.post<Lab3c.CriteriaEstimationResult>(
      '/lab3c/variant/estimation',
      data
    )
  ),
  sendArea: useServiceAction((data: Lab3c.Area) =>
    $apiStudentLab.post<Lab3c.AreaResult>('/lab3c/variant/area', data)
  ),
  sendLine: useServiceAction((data: Lab3c.LineParameters) =>
    $apiStudentLab.post<Lab.CoffAnswer>('/lab3c/variant/line', data)
  ),
  sendQuadratic: useServiceAction((data: Lab3c.KvLine) =>
    $apiStudentLab.post<Lab3c.KvLineResult>('/lab3c/variant/quadratic', data)
  ),
  sendResult: useServiceAction((data: Lab3c.Result) =>
    $apiStudentLab.post<Lab3c.Result>('/lab3c/variant/result', data)
  ),
  getInfo: useServiceAction(() =>
    $apiStudentLab.get<Lab.Info>('/lab3c/variant/info')
  ),
  increment0Step: useServiceAction(() =>
    $apiStudentLab.post<Lab.Info>('/lab3c/variant/increment-zero-step')
  ),
  increment2Step: useServiceAction(() =>
    $apiStudentLab.post<Lab.Info>('/lab3c/variant/increment-second-step')
  ),
};
