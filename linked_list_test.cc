#include <any>
#include <iostream>
#include <ranges>
#include <string>
using Any = std::any;

class LinkNode {
 public:
  Any value;
  LinkNode* next;

 public:
  LinkNode(Any val) {
    value = val;
    next = nullptr;
  }

  void print() {
    try {
      if (value.type() == typeid(int)) {
        std::cout << std::to_string(std::any_cast<int>(value));
        return;
      } else if (value.type() == typeid(double)) {
        std::cout << std::to_string(std::any_cast<double>(value));
        return;
      } else if (value.type() == typeid(std::string)) {
        std::cout << std::any_cast<std::string>(value);
        return;
      }
    } catch (const std::bad_any_cast&) {
      std::cout << "Unsupported type";
      return;
    }
    std::cout << "Unknown type";
  }
};

class LinkedList {
 private:
  LinkNode* head;
  LinkNode* tail;
  int length;

 public:
  LinkedList() {
    head = nullptr;
    tail = nullptr;
    length = 0;
  }
  ~LinkedList() {
    LinkNode* current = head;
    while (current != nullptr) {
      LinkNode* next = current->next;
      delete current;
      current = next;
    }
  }
  void printList() {
    if (length == 0) {
      std::cout << "list is empty" << std::endl;
    };
    LinkNode* current = head;

    while (current != nullptr) {
      current->print();
      current = current->next;
    };
    std::cout << std::endl;
  }

  void append(Any val) {
    LinkNode* node = new LinkNode(val);
    if (length == 0) {
      head = node;
      tail = node;
      length++;
      return;
    }
    tail->next = node;
    tail = node;
    length++;
    return;
  }

  LinkNode* pop() {
    if (length == 0) {
      return nullptr;
    }

    LinkNode* pre = head;
    LinkNode* temp = head;

    while (temp->next != nullptr) {
      pre = temp;
      temp = temp->next;
    }
    if (tail == temp) {
      tail = pre;
      if (tail != nullptr) {
        tail->next = nullptr;
      }
    }
    length--;
    if (length == 0) {
      head = nullptr;
      tail = nullptr;
    }
    return temp;
  }

  void prepend(Any val) {
    LinkNode* node = new LinkNode(val);
    if (length == 0) {
      head = node;
      tail = node;
    } else {
      node->next = head;
      head = node;
    }
    length++;
  }

  LinkNode* popfirst() {
    if (length == 0) {
      return nullptr;
    }

    auto temp = head;
    head = temp->next;

    if (temp) {
      temp->next = nullptr;
    }
    length--;
    if (length == 0) {
      head = nullptr;
      tail = nullptr;
    }
    return temp;
  }

  LinkNode* get(int index) {
    if (index < 0 || index >= length) {
      return nullptr;
    }
    auto temp = head;
    for (int i : std::ranges::iota_view(0, index)) {
      temp = temp->next;
    }
    return temp;
  }

  bool set(unsigned int index, Any val) {
    auto temp = get(index);
    if (temp == nullptr) {
      return false;
    }
    temp->value = val;
    return true;
  }

  bool insert(unsigned int index, Any val) {
    if (length == index) {
      append(val);
      return true;
    } else if (index == 0) {
      prepend(val);
      return true;
    }
    auto node = new LinkNode(val);
    auto temp = get(index - 1);
    if (temp == nullptr) {
      delete node;
      return false;
    }
    node->next = temp->next;
    temp->next = node;
    length++;
    return true;
  }

  LinkNode* remove(unsigned int index) {
    if (length - 1 == index) {
      return pop();
    } else if (index == 0) {
      return popfirst();
    }
    auto temp = get(index - 1);
    auto node = temp->next;
    temp->next = node->next;
    node->next = nullptr;
    length--;
    return node;
  }

  void reverse() {
    if (length == 0 || length == 1) {
      return;
    }

    auto temp = head;
    head = tail;
    tail = temp;

    LinkNode* after = nullptr;
    LinkNode* before = nullptr;
    for (int _ : std::ranges::iota_view(0, length)) {
      after = temp->next;
      temp->next = before;
      before = temp;
      temp = after;
    }
  }
};

int main() {
  LinkedList list;
  list.append(1);
  list.append(2);
  list.printList();
  LinkNode* popnode = list.pop();
  std::cout << "pop : " << popnode << std::endl;

  if (popnode) {
    popnode->print();
    std::cout << std::endl;
    delete popnode;
    popnode = nullptr;
  }
  list.prepend(1);
  list.prepend(std::string("asdsd"));
  popnode = list.popfirst();
  std::cout << "popfirst : " << popnode << std::endl;
  if (popnode) {
    popnode->print();
    std::cout << std::endl;
    delete popnode;
    popnode = nullptr;
  }
  list.printList();

  popnode = list.get(0);
  if (popnode) {
    popnode->print();
    std::cout << std::endl;
    popnode = nullptr;
  }
  std::cout << "out range get" << std::endl;
  popnode = list.get(120);
  std::cout << (popnode == nullptr) << std::endl;
  if (popnode != nullptr) {
    popnode->print();
    std::cout << std::endl;
    popnode = nullptr;
  } else {
    // nullptr
    std::cout << "nullptr" << std::endl;
  }
  list.printList();
  std::cout << list.set(0, std::string("first")) << std::endl;
  list.printList();
  std::cout << list.insert(1, std::string("reverse first")) << std::endl;
  list.printList();
  std::cout << list.remove(2) << std::endl;
  list.printList();
  list.reverse();
  list.printList();
  return 0;
}