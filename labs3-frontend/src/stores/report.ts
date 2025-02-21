import { defineStore } from 'pinia';
import { ReportService } from '../services/report';

export const useReportStore = defineStore('report', () => {
  const getReport = async (
    groupId: number,
    disciplineId: number,
    isExam: boolean
  ) => {
    // await ReportService.getGroupDiscipline({
    //   group_id: groupId,
    //   discipline_id: disciplineId,
    //   is_exam: isExam,
    // });
  };

  return { getReport };
});
