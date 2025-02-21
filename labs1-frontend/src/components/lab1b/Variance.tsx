import React, {useEffect, useState} from "react";
import {useGetVariantQuery} from "../../api/lab1BApi.ts";
import "./styles.scss"
import Timer from "../timer/Timer.tsx";
import AddCriteria from "./addcriteria/AddCriteria.tsx";
import AddAlternatives from "./addalternatives/AddAlternatives.tsx";
import {Alert, Divider, message, Spin, Steps} from "antd";
import {useDispatch, useSelector} from "react-redux";
import {addCompletedStep, getActiveStep, getCompletedSteps, setActiveStep} from "../../reducers/stepSlice.ts";
import PairwiseComparisonCmatricesLvl1 from "./pairwisecomparisonmatrices/lvl1/PairwiseComparisonCmatricesLvl1.tsx";
import PairwiseComparisonCmatricesLvl2 from "./pairwisecomparisonmatrices/lvl2/PairwiseComparisonCmatricesLvl2.tsx";
import EvaluationHierarchyConsistency from "./evaluationhierarchyconsistency/EvaluationHierarchyConsistency.tsx";
import ProcessingCountCriteria from "./processingcountcriteria/ProcessingCountCriteria.tsx";
import ProcessingQualitativeCriteria from "./processingqualitativecriteria/ProcessingQualitativeCriteria.tsx";
import Priorities from "./priorities/Priorities.tsx";
import Results from "./results/Results.tsx";
import GlobalTask from "./task/GlobalTask.tsx";
import ImportanceTable from "./static/ImportanceTable.tsx";
import MatrixSizeTable from "./static/MatrixSizeTable.tsx";
import TaskDescription from "./task/TaskDescription.tsx";
import descriptionsData from "./task/taskDesctiptionsData.ts";

const { Step } = Steps;

const steps = ["Альтернативы", "Критерии", "Шаг 1", "Шаг 2", "Шаг 3", "Шаг 4", "Шаг 5", "Приоритеты", "Результаты"];


const Variance: React.FC = () => {
    const dispatch = useDispatch();
    const activeStep = useSelector(getActiveStep);
    const completedSteps = useSelector(getCompletedSteps);

    const {
        data,
        isLoading,
        error,
        // refetch
    } = useGetVariantQuery({laboratoryId: 2, minutesDuration: 90})


    const changeStep = (newStep: number) => {
        if (completedSteps.includes(newStep - 1) || newStep === 0) {
            dispatch(setActiveStep(newStep));
        } else {
            message.error("Пожалуйста, завершите предыдущий шаг.");
        }
    };

    const completeStep = (stepToCompete: number, notChangeSlide?: boolean) => () => {
        dispatch(addCompletedStep(stepToCompete));

        if (notChangeSlide) {
            return;
        }

        if (activeStep < steps.length - 1) {
            dispatch(setActiveStep(stepToCompete + 1));
        } else {
            message.success("Все шаги завершены!");
        }
    };


    if (isLoading) {
        return (
            <div style={{ textAlign: "center", marginTop: "20px" }}>
                <Spin size="large" />
                <div>Загрузка...</div>
            </div>
        );
    }

    if (error) {
        return (
            <div style={{ marginTop: "20px" }}>
                <Alert
                    message="Ошибка"
                    description={error?.data?.message || "Произошла ошибка!"}
                    type="error"
                    showIcon
                />
            </div>
        );
    }


    const renderStepContent = () => {
        switch (activeStep) {
            case 0:
                return (<>
                    <AddAlternatives onCompleteStep={completeStep(0)} />
                </>)
            case 1:
                return <AddCriteria onCompleteStep={completeStep(1)}/>
            case 2:
                return (
                    <>
                        <ImportanceTable/>
                        <MatrixSizeTable/>
                        <PairwiseComparisonCmatricesLvl1 onCompleteStep={completeStep(2)}/>
                    </>
                )
            case 3:
                return (
                    <>
                        <ImportanceTable/>
                        <MatrixSizeTable/>
                        <PairwiseComparisonCmatricesLvl2 onCompleteStep={completeStep(3, true)}/>
                    </>
                )
            case 4:
                return (
                    <>
                        <MatrixSizeTable/>
                        <EvaluationHierarchyConsistency onCompleteStep={completeStep(4)}/>
                    </>
                )
            case 5:
                return <ProcessingCountCriteria onCompleteStep={completeStep(5)}/>
            case 6:
                return (
                    <>
                        <ImportanceTable/>
                        <MatrixSizeTable/>
                        <ProcessingQualitativeCriteria onCompleteStep={completeStep(6, true)}/>
                    </>
            )
            case 7:
                return <Priorities onCompleteStep={completeStep(7)}/>
            case 8:
                return <Results/>
        }
    }

    return (
        <div className={"container"}>
            <h1 className={"title"}>Лабораторная работ №1Б "Многоуровневые иерархии"</h1>
            <Timer/>
            {data && (
                <>
                    <Steps current={activeStep}
                           onChange={changeStep}
                           direction="horizontal"
                           style={{marginTop: 50}}
                    >
                        {steps.map((step, index) => (
                            <Step key={index} title={step} />
                        ))}
                    </Steps>
                    {activeStep === 0 ? <>
                        <GlobalTask data={JSON.parse(data.variant.variance.variance.task).task}/>
                        <Divider />
                    </> : null}
                    {activeStep !== 8 ? (
                        <TaskDescription taskName={descriptionsData[activeStep]}/>
                    ) : null}
                    {renderStepContent()}
                </>
            )}
        </div>
    )

}

export default Variance