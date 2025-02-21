import { createSlice, PayloadAction } from "@reduxjs/toolkit";

const localStorageName = "criterias"


interface CriteriasSliceProps {
    // criteria: string;
    // extra: {
    //     criteria: string;
    //     is_count: boolean;
    //     value: number[];
    //     is_reverse: boolean;
    // }[];
    criteria: any
}

interface CriteriasState {
    criterias: CriteriasSliceProps;
}

const initialState: CriteriasState = {
    criterias: JSON.parse(localStorage.getItem(localStorageName) || "[]"),
};

const criteriasSlice = createSlice({
    name: "criterias",
    initialState,
    reducers: {
        setCriterias: (state, action: PayloadAction<any>) => {
            state.criterias = action.payload;
            localStorage.setItem(localStorageName, JSON.stringify(state.criterias));
        },
    },
});


export const getCriterias = (state: CriteriasSliceProps[]) => state.criterias.criterias;
export const getMainCriterias = (state: CriteriasSliceProps[]) => state.criterias.criterias.map(item => item.criteria);


export const {
    setCriterias,
} = criteriasSlice.actions;
export default criteriasSlice.reducer;
