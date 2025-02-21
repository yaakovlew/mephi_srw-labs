import {createSlice} from "@reduxjs/toolkit";

const evaluationTableDataLocalStorage = "evaluationTableDataLocalStorage"
const evaluationSummaryDataLocalStorage = "evaluationSummaryDataLocalStorage"
const evaluationResponseResultLocalStorage = "evaluationResponseResultLocalStorage"

export interface evaluationSliceInterface {
    tableData: any;
    summaryData: any;
    responseResult: any;
}


const initialState: evaluationSliceInterface = {
    tableData: JSON.parse(localStorage.getItem(evaluationTableDataLocalStorage) as string),
    summaryData: JSON.parse(localStorage.getItem(evaluationSummaryDataLocalStorage) as string),
    responseResult: JSON.parse(localStorage.getItem(evaluationResponseResultLocalStorage) as string),
}

const evaluationSlice = createSlice({
    name: "evaluation",
    initialState,
    reducers: {
        setTableData: (state, action) => {
            state.tableData = action.payload;
            localStorage.setItem(evaluationTableDataLocalStorage, JSON.stringify(state.tableData));
        },
        setSummaryData: (state, action) => {
            state.summaryData = action.payload;
            localStorage.setItem(evaluationSummaryDataLocalStorage, JSON.stringify(state.summaryData));
        },
        setResponseResult: (state, action) => {
            state.responseResult = action.payload;
            localStorage.setItem(evaluationResponseResultLocalStorage, JSON.stringify(state.responseResult));
        },
    }
})

export const getTableDataEvaluation = (state: evaluationSliceInterface) => state.evaluation.tableData;
export const getSummaryDataEvaluation = (state: evaluationSliceInterface) => state.evaluation.summaryData;
export const getResponseResultEvaluation = (state: evaluationSliceInterface) => state.evaluation.responseResult;

export const {
    setTableData,
    setSummaryData,
    setResponseResult
} = evaluationSlice.actions;

export default evaluationSlice.reducer;