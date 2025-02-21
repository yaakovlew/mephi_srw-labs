import { data } from 'autoprefixer';
import { defineStore } from 'pinia';
import { IDiscipline } from 'src/models/discipline/discipline';
import { IGroup } from 'src/models/group/group';
import { DisciplineService } from 'src/services/discipline';
import { Ref, ref } from 'vue';
import { ISeminarian } from '../models/seminarian/seminarian';

export const useDisciplinesStore = defineStore('disciplines', () => {
  const disciplines: Ref<IDiscipline[] | null> = ref(null);
  const disciplineGroups: Ref<IGroup[] | null> = ref(null);
  const discipline: Ref<IDiscipline.DisciplineFullInfo | null> = ref(null);
  const groups: Ref<IGroup[] | null> = ref(null);
  const avalibaleGroups: Ref<IGroup[] | null> = ref(null);
  const seminarians: Ref<ISeminarian[] | null> = ref(null);

  const getDisciplines = async () => {
    const res = await DisciplineService.fetch();
    if (res.data) {
      disciplines.value = res.data?.ru;
    }
  };

  const getDisciplineGroups = async (id: string) => {
    const res = await DisciplineService.getAllGroups({ id });
    if (res.data) {
      disciplineGroups.value = res.data.groups;
    }
  };

  const addGroupToDiscipline = async (
    discipline_id: string,
    group_id: string
  ) => {
    await DisciplineService.addGroupToDiscipline({
      discipline_id: Number(discipline_id),
      group_id: Number(group_id),
    });
  };

  const deleteGroupFromDiscipline = async (
    discipline_id: string,
    group_id: string
  ) => {
    await DisciplineService.deleteGroupFromDiscipline({
      discipline_id: Number(discipline_id),
      group_id: Number(group_id),
    });
  };

  const getAllGroups = async (discipline_id: number) => {
    const res = await DisciplineService.getAllGroups({
      id: String(discipline_id),
    });
    if (res.data) {
      groups.value = res.data.groups;
    }
  };

  const getAvalibaleGroups = async (discipline_id: number) => {
    const res = await DisciplineService.getAlvaliableGroups({
      id: String(discipline_id),
    });
    if (res.data) {
      avalibaleGroups.value = res.data.groups;
    }
  };

  const getDisciplineInfo = async (discipline_id: string) => {
    const res = await DisciplineService.getDisciplineInfo({
      id: discipline_id,
    });
    if (res.data) {
      discipline.value = res.data;
    }
  };

  const createDiscipline = async (discipline: IDiscipline.CreateDiscipline) => {
    await DisciplineService.createDiscipline(discipline);
  };

  const changeDisciplineName = async (
    data: IDiscipline.ChangeDisciplineName
  ) => {
    await DisciplineService.changeDisciplineName(data);
  };

  const changeDisciplineExamMark = async (
    data: IDiscipline.ChangeDisciplineExamMark
  ) => {
    await DisciplineService.changeDisciplineExamMark(data);
  };

  const changeDisciplineLessonMark = async (
    data: IDiscipline.ChangeDisciplineLessonMark
  ) => {
    await DisciplineService.changeDisciplineLessonMark(data);
  };

  const changeDisciplineSeminarMark = async (
    data: IDiscipline.ChangeDisciplineSeminarMark
  ) => {
    await DisciplineService.changeDisciplineSeminarMark(data);
  };

  const getGroupSeminarians = async (disciplineId: string, groupId: string) => {
    const res = await DisciplineService.getGroupSeminarians({
      discipline_id: Number(disciplineId),
      group_id: Number(groupId),
    });
    if (res.data) {
      seminarians.value = res.data.seminarians;
    }
  };

  const addSeminarianToGroup = async (data: IGroup.AddSeminarian) => {
    await DisciplineService.addSeminarianToGroup(data);
  };

  return {
    disciplines,
    disciplineGroups,
    discipline,
    groups,
    avalibaleGroups,
    getDisciplines,
    getDisciplineGroups,
    changeDisciplineName,
    addGroupToDiscipline,
    deleteGroupFromDiscipline,
    getAllGroups,
    getAvalibaleGroups,
    getDisciplineInfo,
    createDiscipline,
    changeDisciplineExamMark,
    changeDisciplineLessonMark,
    changeDisciplineSeminarMark,
    getGroupSeminarians,
    seminarians,
    addSeminarianToGroup,
  };
});
