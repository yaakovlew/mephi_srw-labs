import { defineStore } from 'pinia';
import { IDiscipline } from 'src/models/discipline/discipline';
import { IGroup } from 'src/models/group/group';
import { ISeminarian } from 'src/models/seminarian/seminarian';
import { GroupsService } from 'src/services/groups';
import { IFutureResource } from 'src/utils/service/action';
import { Ref, ref } from 'vue';
import { TestService } from '../services/test';
import { useTestsStore } from 'src/stores/test';

export const useGroupsStore = defineStore('groups', () => {
  const groups: Ref<IGroup[] | null> = ref(null);
  const groupStudents: Ref<IGroup.GroupStudentStudentId[] | null> = ref(null);
  const groupDisciplines: Ref<IDiscipline[] | null> = ref(null);
  const commonGroups: Ref<IGroup[] | null> = ref(null);
  const seminarians: Ref<ISeminarian.Seminarian[] | null> = ref(null);
  const testStore = useTestsStore();

  const getGroups = async () => {
    const res = await GroupsService.fetch();
    if (res.data) {
      groups.value = res.data.groups;
    }
  };

  const getCommonGroups = async () => {
    const res = await GroupsService.fetchCommon();
    if (res.data) {
      commonGroups.value = res.data.groups;
    }
  };

  const addNewGroup = async (name: string) => {
    await GroupsService.addNewGroup({
      group_name: name,
    });
  };

  const changeGroupName = async (name: string, id: number) => {
    await GroupsService.changeGroupName({
      name,
      group_id: id,
    });
  };

  const deleteGroup = async (id: string) => {
    await GroupsService.deleteGroup({
      id,
    });
  };

  const getGroupStudents = async (id: string) => {
    groupStudents.value = null;
    const res = await GroupsService.getGroupStudents({
      id,
    });
    if (res.data?.students) {
      groupStudents.value = res.data?.students;
    }
  };

  const getGroupDiscipline = async (id: string) => {
    groupDisciplines.value = null;
    const res = await GroupsService.getGroupDiscipline({
      id,
    });
    if (res.data?.ru) {
      groupDisciplines.value = res.data?.ru;
    }
  };

  const getSeminarians = async (disciplineId: string, groupId: string) => {
    await GroupsService.getGroupSeminarians({
      discipline_id: Number(disciplineId),
      group_id: Number(groupId),
    });
  };

  const openTestForGroup = async (
    groupId: number,
    testId: number,
    date: number
  ) => {
    const res = await GroupsService.getGroupStudents({
      id: String(groupId),
    });
    if (res.data?.students) {
      const students = res.data.students;
      const openPromise: Promise<void>[] = [];
      students.forEach((student) => {
        openPromise.push(
          testStore.openTest({
            user_id: Number(student.student_id),
            test_id: testId,
            date,
          })
        );
      });
      await Promise.all(openPromise);
    }
  };

  const openTestForStudents = async (
    groupId: number[],
    testId: number,
    date: number
  ) => {
    const openPromise: Promise<void>[] = [];
    groupId.forEach((id) => {
      openPromise.push(
        testStore.openTest({
          user_id: Number(id),
          test_id: testId,
          date,
        })
      );
    });
    await Promise.all(openPromise);
  };

  const closeTestForStudents = async (
    groupId: number[],
    testId: number,
  ) => {
    const openPromise: Promise<void>[] = [];
    groupId.forEach((id) => {
      openPromise.push(
        testStore.closeTest({
          user_id: Number(id),
          test_id: testId,
        })
      );
    });
    await Promise.all(openPromise);
  };

  return {
    groups,
    groupStudents,
    groupDisciplines,
    commonGroups,
    getGroups,
    addNewGroup,
    changeGroupName,
    deleteGroup,
    getGroupStudents,
    getGroupDiscipline,
    getCommonGroups,
    getSeminarians,
    openTestForGroup,
    openTestForStudents,
    closeTestForStudents,
  };
});
