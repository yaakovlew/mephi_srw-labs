import { $apiLecturer } from 'src/boot/axios';
import { ISeminarian } from 'src/models/seminarian/seminarian';
import { useServiceAction } from 'src/utils/service/action';

export const SeminariansService = {
  fetch: useServiceAction(() =>
    $apiLecturer.get<ISeminarian.FetchData>('/seminarian')
  ),
  addToDiscipline: useServiceAction((data: ISeminarian.GroupSeminarian) =>
    $apiLecturer.post('/seminarian', data)
  ),
};
