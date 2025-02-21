import { $apiLecturer, $apiStudent } from 'src/boot/axios';
import { IMark } from 'src/models/mark/mark';
import { useServiceAction } from 'src/utils/service/action';

export const MarksService = {
  getExamMark: useServiceAction((data: IMark.GetMarkExam) =>
    $apiLecturer.get<IMark.Marks>(
      `/mark/exam?group_id=${data.group_id}&discipline_id=${data.discipline_id}`
    )
  ),
  getTestMark: useServiceAction((data: IMark.GetMarkTest) =>
    $apiLecturer.get<IMark.Marks>(
      `/mark/test?group_id=${data.group_id}&test_id=${data.test_id}`
    )
  ),
  getLabaratoryMark: useServiceAction((data: IMark.GetMarkLabaratory) =>
    $apiLecturer.get<IMark.Marks>(
      `/mark/labaratory?group_id=${data.group_id}&labaratory_id=${data.labaratory_id}`
    )
  ),
  postLabaratoryMark: useServiceAction((data: IMark.PostMarkLabaratory) =>
    $apiLecturer.post('/mark/labaratory', data)
  ),
  postTestMark: useServiceAction((data: IMark.PostMarkTest) =>
    $apiLecturer.put('/mark/test', data)
  ),
  postExamMark: useServiceAction((data: IMark.PostMarkExam) =>
    $apiLecturer.post('/mark/exam', data)
  ),
  getTestsMarks: useServiceAction((id: number) =>
    $apiStudent.get<IMark.Marks>(`/marks/test/${id}`)
  ),
  getTestMarkStudent: useServiceAction((id: number) =>
    $apiStudent.get<IMark.TestMark>(`/test/mark/${id}`)
  ),
};
