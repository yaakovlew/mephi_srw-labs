import {createSlice} from "@reduxjs/toolkit";

const qualitativeMatrixLocalStorage = "qualitativeMatrixLocalStorage"
const qualitativeDataLocalStorage = "qualitativeDataLocalStorage"
const qualitativeResponseResultLocalStorage = "qualitativeResponseResultLocalStorage"

export interface qualitativeSliceInterface {
    matrix: string[][][];
    data: string[][];
    responseResult: string[][];
}

const initialState: qualitativeSliceInterface = {
    matrix: JSON.parse(localStorage.getItem(qualitativeMatrixLocalStorage) as string) || [],
    data: JSON.parse(localStorage.getItem(qualitativeDataLocalStorage) as string) || [],
    responseResult: JSON.parse(localStorage.getItem(qualitativeResponseResultLocalStorage) as string) || [],
}


const qualitativeSlice = createSlice({
    name: "qualitative",
    initialState,
    reducers: {
        setMatrix: (state, action) =>  {
            const {payload, index} = action.payload;
            state.matrix[index] = payload;
            localStorage.setItem(qualitativeMatrixLocalStorage, JSON.stringify(state.matrix));
        },
        setData: (state, action) => {
            const {payload, index} = action.payload;
            state.data[index] = payload;
            localStorage.setItem(qualitativeDataLocalStorage, JSON.stringify(state.data));
        },
        setResponseResult: (state, action) => {
            const {payload, index} = action.payload;
            state.responseResult[index] = payload;
            localStorage.setItem(qualitativeResponseResultLocalStorage, JSON.stringify(state.responseResult));
        },
    }
})


export const getMatrixQualitative = (index) => (state: qualitativeSliceInterface) => state.qualitative.matrix[index] || null;
export const getDataQualitative = (index) => (state: qualitativeSliceInterface) => state.qualitative.data[index] || null;
export const getResponseResultQualitative = (index) => (state: qualitativeSliceInterface) => state.qualitative.responseResult[index] || null;

export const {
    setMatrix,
    setData,
    setResponseResult,
} = qualitativeSlice.actions;

export default qualitativeSlice.reducer;