import React, {useEffect, useState} from "react";
import {List, Input, Button, Space, Typography, Divider, message} from "antd";
import { PlusOutlined, DeleteOutlined } from "@ant-design/icons";
import {useAddAlternativesMutation} from "../../../api/lab1BApi.ts";
import {useDispatch, useSelector} from "react-redux";
import {getAlternatives, setAlternatives as setAlternativesDispatch} from "../../../reducers/alternativesSlice.ts";

const { Title } = Typography;


interface Props {
    onCompleteStep: () => void
}

const validateAlternatives = (alternatives: string[]) => {
	if (alternatives.length < 3 || alternatives.length > 5) {
			return "Количество альтернатив должно быть от 3 до 5.";
	}

	return null;
};

const AddAlternatives: React.FC<Props> = ({onCompleteStep}) => {
    const dispatch = useDispatch();
    const initialAlternatives = useSelector(getAlternatives);
    const [alternatives, setAlternatives] = useState<string[]>([]);
    const [inputValue, setInputValue] = useState<string>("");
    const [sendAlternatives] = useAddAlternativesMutation();
    const [disabledInput, setDisabledInput] =  useState(false)


    useEffect(() => {
        if (initialAlternatives && initialAlternatives.length > 0) {
            setAlternatives(initialAlternatives);
            setDisabledInput(true)
        }
    }, [initialAlternatives]);

    const addAlternative = () => {
        if (inputValue.trim()) {
            setAlternatives([...alternatives, inputValue.trim()]);
            setInputValue("");
        }
    };

    const removeAlternative = (index: number) => {
        setAlternatives(alternatives.filter((_, i) => i !== index));
    };

    const onSend = async () => {
				const validation = validateAlternatives(alternatives)
				
				if (validation) {
						message.error(validation)
						return;
				}

        try {
            const response = await sendAlternatives({ alternatives: alternatives });

            if (response.error) {
                // Если ошибка (например, статус 401), можно показать сообщение
                message.error("Ошибка: не авторизован или другие проблемы с запросом.");
                return; // Не выполняем дальнейшую логику
            }

            dispatch(setAlternativesDispatch(alternatives));
            onCompleteStep();
            console.log("Сохранить альтернативы", alternatives);
        } catch (error) {
            // Логирование и обработка ошибок
            message.error("Произошла ошибка при сохранении альтернатив.");
            console.error(error);
        }
    };

    const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
        if (e.key === "Enter") {
            addAlternative();
        }
    };

    const renderHeader = ()  => {
        return (
            <Space.Compact style={{ width: '100%' }}>
                <Input
                    disabled={disabledInput}
                    value={inputValue}
                    onChange={(e) => setInputValue(e.target.value)}
                    placeholder="Введите альтернативу"
                    onKeyDown={handleKeyPress}
                />
                <Button
                    disabled={disabledInput}
                    type="primary"
                    icon={<PlusOutlined />}
                    onClick={addAlternative}
                >
                    Добавить альтернативу
                </Button>
            </Space.Compact>
        )
    }

    const renderFooter = () => {
        return (
            <Button
                disabled={disabledInput}
                type="primary"
                onClick={onSend}
                style={{ marginTop: 16 }}
            >
                Сохранить альтернативы
            </Button>

        )
    }

    return (
        <>
            <Title level={4} style={{ textAlign: "center" }}>Альтернативы</Title>
            <List
                header={renderHeader()}
                footer={renderFooter()}
                bordered
                dataSource={alternatives}
                style={{width: 500, margin: "auto"}}
                renderItem={(item, index) => (
                    <List.Item
                        key={index}
                        actions={[
                            <Button
                                disabled={disabledInput}
                                type="link"
                                danger
                                icon={<DeleteOutlined />}
                                onClick={() => removeAlternative(index)}
                            >
                            </Button>,
                        ]}
                    >
                        {item}
                    </List.Item>
                )}
            />
        </>
    );
};


export default AddAlternatives;
