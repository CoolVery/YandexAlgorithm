package main

type LinkedList struct {
	value int
	nextList *LinkedList
}
//Функция 1 запроса - добавить y после х
func (ll *LinkedList) AddYAfterX(x, y int) *LinkedList {
	//Если x=0, то нужно сделать число y новым началом списка
	if x == 0 {
		//Если у нас список nil, значит еще ни разу не создавали голову, т.е. еще ни разу не создавали элемент списка
		if ll == nil {
			return &LinkedList{
				value: y,
				nextList: nil,
			}
		}
		//Иначе мы создаем новый лист и ссылаем на старую голову, т.е. уже в списке уже есть элементы
		return &LinkedList{
			value: y,
			nextList: ll,
		}
	}
	//Сохраняем список head - чтобы вернуть начало списка, currentList - чтобы проходить по нему
	head, currentList := ll, ll
	//Счетчик позиции элементов, начиаем с головы - 1
	countLink := 1
	for {
		//Если счетчик равен искомому, то добавляем (После 3 элемента, например)
		if countLink == x {
			//Здесь мы создаем элемент, который ссылается на следующий элемент, после чего мы меняем у current ссылку на новый элемент
			newList := &LinkedList{
				value: y,
				nextList: currentList.nextList,
			}
			currentList.nextList = newList
			//Возвращаем начало списка
			return head
		}
		//Увеливаем счетчик и переходим в другой элемент
		countLink++
		currentList = currentList.nextList
	}
	
}
//Функция 2 запроса - найти позицию в списке
func (ll *LinkedList) SearchPosithionInList(searchPosithion int) int {
	//Мы знаем, что в списке есть 1 элемент
	countLink := 1
	for {
		//Если позициия и количество пройденных ссылок равны - выводим значение
		if countLink == searchPosithion {
			return ll.value
		}
		//Переходим и увеличиваем значение
		ll = ll.nextList
		countLink++
	}
}
//Функция 3 запроса - удалить число по позиции 
func (ll *LinkedList) DeleteForPosithion(deletePosithion int) *LinkedList {
	//Храанит родителя
	parrentLink := ll 
	//Голова списка
	head := ll
	countLink := 1
	for {
		//Если позиция 1 - это голова, если у него нет ссылки далее, то наш список перестает существовать
		if deletePosithion == 1 && ll.nextList == nil {
			return nil
		} 
		//Если позиция 1 - это голова, но у него дальше есть ссылка, значит мы голову делаем тот элемент, на который ссылается голова
		if deletePosithion == 1 {
			return ll.nextList
		}
		//Если мы дошли до позиции, то мы в ссылку след элемента родителя кладем ссылку на след элемента конкретного элемента
		if countLink == deletePosithion {
			parrentLink.nextList = ll.nextList
			return head
		}
		//Переходим и увеличиваем значение
		parrentLink = ll 
		ll = ll.nextList
		countLink++
	}
}