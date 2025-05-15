package main

/*
### Состояния процесса создания карточки товара:
1. **Draft** (Черновик) - начальное состояние, когда карточка товара создается.
2. **Review** (На проверке) - карточка отправлена на проверку.
3. **Approved** (Одобрено) - карточка товара одобрена и готова к публикации.
4. **Published** (Опубликовано) - карточка товара опубликована на маркетплейсе.
5. **Rejected** (Отклонено) - карточка товара отклонена.

### Переходы между состояниями:
- Из состояния **Draft** можно перейти в **Review**.
- Из состояния **Review** можно перейти в **Approved** или **Rejected**.
- Из состояния **Approved** можно перейти в **Published**.

### Пример кода:

*/

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	ctx := context.Background()
	// Создаем новый конечный автомат
	fsm := fsm.NewFSM(
		"draft", // начальное состояние
		fsm.Events{
			{Name: "submit", Src: []string{"draft"}, Dst: "review"},
			{Name: "approve", Src: []string{"review"}, Dst: "approved"},
			{Name: "reject", Src: []string{"review"}, Dst: "rejected"},
			{Name: "publish", Src: []string{"approved"}, Dst: "published"},
		},
		fsm.Callbacks{
			"enter_state": func(ctx context.Context, e *fsm.Event) {
				fmt.Println("Entered state:", e.Dst)
			},
			"leave_state": func(ctx context.Context, e *fsm.Event) {
				fmt.Println("Leaving state:", e.Src)
			},
		},
	)

	// Начальное состояние
	fmt.Println("Current state:", fsm.Current())

	// Переход из Draft в Review
	err := fsm.Event(ctx, "submit")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current state:", fsm.Current())

	// Переход из Review в Approved
	err = fsm.Event(ctx, "approve")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current state:", fsm.Current())

	// Переход из Approved в Published
	err = fsm.Event(ctx, "publish")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current state:", fsm.Current())

	// Попробуем отклонить уже одобренный товар (должно вернуть ошибку)
	err = fsm.Event(ctx, "reject")
	if err != nil {
		fmt.Println("err : ", err)
	}
}

/*
### Описание кода:
1. Мы создаем новый конечный автомат с начальными состоянием `draft`.
2. Определяем события, которые могут произойти, и соответствующие переходы между состояниями.
3. Используем колбэки для вывода информации о входе и выходе из состояний.
4. В конце мы демонстрируем последовательность переходов от состояния `draft` до `published`.
Этот пример можно расширять, добавляя дополнительные состояния и события по мере необходимости, а также обрабатывать ошибки более детально.
*/
