#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// 定义一个字符串类型
typedef struct string {
    char *data;   // 指向字符串数据的指针
    int length;   // 字符串的长度
} String;

// 定义链表节点类型
typedef struct node {
    String *string;  // 指向字符串的指针
    struct node *next; // 指向下一个节点的指针
} Node;

// 创建一个字符串
String *create_string(char *data) {
    String *string = (String *)malloc(sizeof(String));
    string->length = strlen(data);
    string->data = (char *)malloc(sizeof(char) * (string->length + 1));
    strcpy(string->data, data);
    return string;
}

// 销毁一个字符串
void destroy_string(String *string) {
    free(string->data);
    free(string);
}

// 创建一个链表节点
Node *create_node(String *string) {
    Node *node = (Node *)malloc(sizeof(Node));
    node->string = string;
    node->next = NULL;
    return node;
}

// 销毁一个链表节点
void destroy_node(Node *node) {
    destroy_string(node->string);
    free(node);
}

// 在链表末尾插入一个节点
void append_node(Node **head, Node *node) {
    if (*head == NULL) {
        *head = node;
    } else {
        Node *current = *head;
        while (current->next != NULL) {
            current = current->next;
        }
        current->next = node;
    }
}

// 打印链表中的所有字符串
void print_list(Node *head) {
    Node *current = head;
    while (current != NULL) {
        printf("%s\n", current->string->data);
        current = current->next;
    }
}

// 连接两个字符串
String *concatenate_strings(String *s1, String *s2) {
    char *data = (char *)malloc(sizeof(char) * (s1->length + s2->length + 1));
    strcpy(data, s1->data);
    strcat(data, s2->data);
    String *string = create_string(data);
    free(data);
    return string;
}

// 反转字符串
String *reverse_string(String *string) {
    char *data = (char *)malloc(sizeof(char) * (string->length + 1));
    int i;
    for (i = 0; i < string->length; i++) {
        data[i] = string->data[string->length - i - 1];
    }
    data[string->length] = '\0';
    String *new_string = create_string(data);
    free(data);
    return new_string;
}

// 主函数
int main() {
    // 创建链表
    Node *head = NULL;

    // 输入字符串
    char data[100];
    printf("Please input a string: ");
    scanf("%s", data);
    String *string1 = create_string(data);
    Node *node1 = create_node(string1);
    append_node(&head, node1);

    // 输入字符串
    printf("Please input another string: ");
    scanf("%s", data);
