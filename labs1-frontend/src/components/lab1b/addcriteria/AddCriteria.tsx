import React, {useEffect, useState} from "react";
import {Card, List, Input, Button, Space, Checkbox, Typography, Divider, message} from "antd";
import { PlusOutlined, DeleteOutlined } from "@ant-design/icons";
import { useAddCriteriaMutation } from "../../../api/lab1BApi.ts";
import {useDispatch, useSelector} from "react-redux";
import { getAlternativesNumber } from "../../../reducers/alternativesSlice.ts";
import "./addcriteria.scss"
import {getCriterias, setCriterias} from "../../../reducers/criteriasSlice.ts";

const { Title } = Typography;

interface Subsection {
    id: number;
    name: string;
    isCount: boolean;
    isReverse: boolean;
    value: number[];
}

interface Section {
    id: number;
    criteria: string;
    extra: Subsection[];
}

interface Props {
    onCompleteStep: () => void
}

const AddCriteria: React.FC<Props> = ({onCompleteStep}) => {
    const dispatch = useDispatch();
    const [sections, setSections] = useState<Section[]>([]);
    const [idCounter, setIdCounter] = useState<number>(1);
    const numberAlternatives = useSelector(getAlternativesNumber);
    const initialSections = useSelector(getCriterias);
    const [disabledInput, setDisabledInput] =  useState(false)
    const [sendCriterias, {
        error,
        isSuccess,
        isLoading
    }] = useAddCriteriaMutation();

    useEffect(() => {
        if (initialSections && initialSections.length > 0) {
            setSections(initialSections);
            setDisabledInput(true)
        }
    }, [initialSections]);

    useEffect(() => {
        if (error) {
            console.error(error);
            message.error(JSON.stringify(error));
        }
    }, [error]);

    useEffect(() => {
        if (isSuccess) {
            dispatch(setCriterias(sections))
            onCompleteStep();
        }
    }, [dispatch, isSuccess, onCompleteStep, sections]);

    const addSection = () => {
        setSections([
            ...sections,
            { id: idCounter, criteria: `Основной критерий ${idCounter}`, extra: [] },
        ]);
        setIdCounter((prev) => prev + 1);
    };

    const addSubsection = (sectionId: number) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId
                    ? {
                        ...section,
                        extra: [
                            ...section.extra,
                            {
                                id: idCounter,
                                name: `Дополнительный критерий ${idCounter}`,
                                isCount: false,
                                isReverse: false,
                                value: new Array(numberAlternatives).fill(0),
                            },
                        ],
                    }
                    : section
            )
        );
        setIdCounter((prev) => prev + 1);
    };

    const updateSectionName = (sectionId: number, newName: string) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId ? { ...section, criteria: newName } : section
            )
        );
    };

    const updateSubsectionName = (sectionId: number, subId: number, newName: string) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId
                    ? {
                        ...section,
                        extra: section.extra.map((sub) =>
                            sub.id === subId ? { ...sub, name: newName } : sub
                        ),
                    }
                    : section
            )
        );
    };

    const toggleIsCount = (sectionId: number, subId: number) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId
                    ? {
                        ...section,
                        extra: section.extra.map((sub) =>
                            sub.id === subId
                                ? {
                                    ...sub,
                                    isCount: !sub.isCount,
                                    value: !sub.isCount
                                        ? new Array(numberAlternatives).fill(0)
                                        : [],
                                }
                                : sub
                        ),
                    }
                    : section
            )
        );
    };

    const toggleIsReverse = (sectionId: number, subId: number) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId
                    ? {
                        ...section,
                        extra: section.extra.map((sub) =>
                            sub.id === subId ? { ...sub, isReverse: !sub.isReverse } : sub
                        ),
                    }
                    : section
            )
        );
    };

    const handleValueChange = (sectionId: number, subId: number, index: number, newValue: number) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId
                    ? {
                        ...section,
                        extra: section.extra.map((sub) =>
                            sub.id === subId
                                ? {
                                    ...sub,
                                    value: sub.value.map((v, i) =>
                                        i === index ? newValue : v
                                    ),
                                }
                                : sub
                        ),
                    }
                    : section
            )
        );
    };

    const removeSection = (sectionId: number) => {
        setSections(sections.filter((section) => section.id !== sectionId));
    };

    const removeSubsection = (sectionId: number, subId: number) => {
        setSections(
            sections.map((section) =>
                section.id === sectionId
                    ? {
                        ...section,
                        extra: section.extra.filter((sub) => sub.id !== subId),
                    }
                    : section
            )
        );
    };

    const validateCriterias = (criterias: { criteria: string; extra: { criteria: string; is_count: boolean; value: any; is_reverse: boolean }[] }[]) => {
        if (criterias.length < 3 || criterias.length > 5) {
            return "Количество критериев должно быть от 3 до 5.";
        }

				let countCountCriteria = 0
				let countQualityCriteria = 0
        for (let i = 0; i < criterias.length; i++) {
            const extra = criterias[i].extra;
            if (extra.length < 3 || extra.length > 10) {
                return "Основной критерий должен содержать от 3 до 10 доп. критериев";
            }
						
						for (let j = 0; j < extra.length; j++) {
							if (extra[j].is_count) {
									countCountCriteria++;
							} else {
									countQualityCriteria++;
							}
					}
        }

				if (countCountCriteria <= 2) {
						return "Количество количественных критериев должно быть больше 2.";
				}

				if (countQualityCriteria <= 2) {
					return "Количество качественных критериев должно быть больше 2.";
			}

      return null;
    };


    const onSend = async () => {
        try {
            const criterias = sections.map((section) => ({
                criteria: section.criteria,
                extra: section.extra.map((sub) => ({
                    criteria: sub.name,
                    is_count: sub.isCount,
                    value: sub.isCount ? sub.value : [],
                    is_reverse: sub.isReverse,
                })),
            }));

            const validation = validateCriterias(criterias)

            if (validation) {
                message.error(validation)
                return;
            }

            const response = await sendCriterias({ criterias });

            if (response.error) {
                message.error("Ошибка: не удалось сохранить критерии.");
                return;
            }
        } catch (error) {
            message.error("Произошла ошибка при отправке данных.");
            console.error(error);
        }
    };


    const renderItem = (sub, section) => (
        <List.Item key={sub.id}>
            <Space direction="vertical" style={{ width: "100%" }}>
                <Space.Compact style={{ width: "100%" }}>
                    <Input
                        readOnly={disabledInput}
                        value={sub.name}
                        onChange={(e) =>
                            updateSubsectionName(section.id, sub.id, e.target.value)
                        }
                        placeholder="Название доп. критерия"
                    />
                    <Button danger
                            icon={<DeleteOutlined />}
                            onClick={() => removeSubsection(section.id, sub.id)}
                            disabled={disabledInput}
                    />
                </Space.Compact>

                <Space>
                    <Checkbox
                        disabled={disabledInput}
                        checked={sub.isCount}
                        onChange={() => toggleIsCount(section.id, sub.id)}
                    >
                        Числовое значение
                    </Checkbox>
                    <Checkbox
                        disabled={disabledInput}
                        checked={sub.isReverse}
                        onChange={() => toggleIsReverse(section.id, sub.id)}
                    >
                        Обратный порядок
                    </Checkbox>
                </Space>

                {sub.isCount && (
                    <Card title={<div style={{fontWeight: 300}}>Значение соответсвующей альтернативы по критерию</div>}
                          size={"small"}
                    >
                        <Space>
                            {sub.value.map((val, index) => (
                                <Input
                                    readOnly={disabledInput}
                                    key={index}
                                    type="number"
                                    value={val}
                                    onChange={(e) =>
                                        handleValueChange(section.id, sub.id, index, Number(e.target.value))
                                    }
                                    style={{ width: 60 }}
                                />
                            ))}
                        </Space>
                    </Card>

                )}
            </Space>
        </List.Item>
    )

    return (
        <Card style={{ margin: "0 auto", padding: 16 }} bordered>
            <Title level={4} style={{ textAlign: "center" }}>Критерии</Title>
            <Divider />

            <Space wrap={true}>
                {sections.map((section) => (
                    <Card key={section.id} style={{ marginBottom: 16 }} bordered>
                        <Space.Compact style={{ width: "100%" }}>
                            <Input
                                readOnly={disabledInput}
                                value={section.criteria}
                                onChange={(e) => updateSectionName(section.id, e.target.value)}
                                placeholder="Основной критерий"
                            />
                            <Button danger
                                    icon={<DeleteOutlined />}
                                    onClick={() => removeSection(section.id)}
                                    disabled={disabledInput}
                            />
                        </Space.Compact>

                        <List
                            bordered
                            dataSource={section.extra}
                            style={{ marginTop: 16 }}
                            renderItem={(sub) => renderItem(sub, section)}
                        />

                        <Button
                            disabled={disabledInput}
                            type="dashed"
                            icon={<PlusOutlined />}
                            onClick={() => addSubsection(section.id)}
                            block
                            style={{ marginTop: 10 }}
                            loading={isLoading}
                        >
                            Добавить дополнительный критерий
                        </Button>
                    </Card>
                ))}
                <Button variant="outlined"
                        icon={<PlusOutlined />}
                        onClick={addSection}
                        block
                        style={{ height: 100}}
                        disabled={disabledInput}
                >
                    Добавить основной критерий
                </Button>
            </Space>

            <Divider />

            <Button type="primary"
                    onClick={onSend}
                    disabled={disabledInput}
                    block
            >
                Сохранить критерии
            </Button>
        </Card>
    );
};

export default AddCriteria;
