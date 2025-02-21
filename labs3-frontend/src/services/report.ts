import { $apiLecturer } from 'src/boot/axios';
import { IReport } from 'src/models/report/report';
import { useServiceAction } from 'src/utils/service/action';

export const ReportService = {
  getGroupDiscipline: useServiceAction((data: IReport.GetReport) =>
    $apiLecturer.get(
      `/group/report?group_id=${data.group_id}&discipline_id=${data.discipline_id}&is_exam=${data.is_exam}`
    )
  ),
};
