import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const localStorageName = "marksData";

interface Mark {
    id: string;
    mark: number;
    maxMark: number;
}

interface MarksState {
    marks: Mark[];
}

const initialState: MarksState = {
    marks: JSON.parse(localStorage.getItem(localStorageName) || "[]"),
};

const marksSlice = createSlice({
    name: "marks",
    initialState,
    reducers: {
        addOrUpdateMark(state, action: PayloadAction<Mark>) {
            const { id, mark, maxMark } = action.payload;
            const existingIndex = state.marks.findIndex((item) => item.id === id);

            if (existingIndex !== -1) {
                state.marks[existingIndex].mark = mark;
                state.marks[existingIndex].maxMark = maxMark;
            } else {
                state.marks.push({ id, mark, maxMark });
            }

            localStorage.setItem(localStorageName, JSON.stringify(state.marks));
        },
    },
});

export const getMarks = (state: any) => state.marks.marks;
export const getMark = (markId: string) => (state: any) => (
    state.marks.marks.find((mark: Mark) => mark.id.includes(markId))
)


export const { addOrUpdateMark } = marksSlice.actions;

export default marksSlice.reducer;
