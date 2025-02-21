import React from "react";

type TaskBarProps = {
    activeStep: number;
    steps: string[];
    onStepChange: (step: number) => void;
};

const TaskBar: React.FC<TaskBarProps> = ({ activeStep, steps, onStepChange }) => {
    return (
        <div style={styles.taskBarContainer}>
            {steps.map((step, index) => (
                <button
                    key={index}
                    style={{
                        ...styles.taskButton,
                        ...(index === activeStep ? styles.activeButton : {}),
                    }}
                    onClick={() => onStepChange(index)}
                >
                    {step}
                </button>
            ))}
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    taskBarContainer: {
        display: "flex",
        justifyContent: "center",
        gap: "10px",
        marginBottom: "20px",
    },
    taskButton: {
        padding: "10px 15px",
        fontSize: "14px",
        border: "1px solid #ccc",
        borderRadius: "5px",
        cursor: "pointer",
        backgroundColor: "#f0f0f0",
        color: "#333",
        transition: "background-color 0.3s ease",
    },
    activeButton: {
        backgroundColor: "#007bff",
        color: "#fff",
        borderColor: "#0056b3",
    },
};

export default TaskBar;