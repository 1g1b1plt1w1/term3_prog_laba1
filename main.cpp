#include <iostream>
#include <string>
#include <map>
#include <vector>
#include <fstream>
#include <sstream>

#include "linked_list.h"
#include "one_linked_list.h"
#include "array.h"
#include "queue.h"
#include "stack.h"
#include "tree.h"

using namespace std;

enum class Command {
    PRINT,

    // --- ARRAY (M — массив) ---
    MCREATE,
    MPUSH,
    MPUSHIDX,
    MGET,
    MDELIDX,
    MSETIDX,
    MLENGTH,
    MREAD,

    // --- ONE-LINKED LIST (F — односвязный список) ---
    FPUSH,
    FPUSHHEAD,
    FPUSHBEFORE,
    FPUSHAFTER,
    FPOP,
    FPOPHEAD,
    FPOPBEFORE,
    FPOPAFTER,
    FDELVAL,
    FFINDVAL,
    FREADHT,
    FREADTH,

    // --- LINKED LIST (L — двусвязный список) ---
    LPUSHHEAD,
    LPUSH,
    LPUSHBEFORE,
    LPUSHAFTER,
    LPOPHEAD,
    LPOP,
    LPOPBEFORE,
    LPOPAFTER,
    LDELVAL,
    LFINDVAL,
    LREADHT,
    LREADTH,

    // --- STACK (S — стек) ---
    SPUSH,
    SPOP,
    SREAD,

    // --- QUEUE (Q — очередь) ---
    QPUSH,
    QPOP,
    QREAD,

    // --- TREE (T — AVL-дерево) ---
    TINSERT,
    TFIND,
    TDELETE,
    TREAD,
    TCHECKAVL
};

map<string, Command> command_map = {
    {"PRINT",Command::PRINT},

    // --- ARRAY ---
    {"MCREATE", Command::MCREATE},
    {"MPUSH", Command::MPUSH},
    {"MPUSHIDX", Command::MPUSHIDX},
    {"MGET", Command::MGET},
    {"MDELIDX", Command::MDELIDX},
    {"MSETIDX", Command::MSETIDX},
    {"MLENGTH", Command::MLENGTH},
    {"MREAD", Command::MREAD},

    // --- ONE-LINKED (F) ---
    {"FPUSH", Command::FPUSH},
    {"FPUSHHEAD", Command::FPUSHHEAD},
    {"FPUSHBEFORE", Command::FPUSHBEFORE},
    {"FPUSHAFTER", Command::FPUSHAFTER},
    {"FPOP", Command::FPOP},
    {"FPOPHEAD", Command::FPOPHEAD},
    {"FPOPBEFORE", Command::FPOPBEFORE},
    {"FPOPAFTER", Command::FPOPAFTER},
    {"FDELVAL", Command::FDELVAL},
    {"FFINDVAL", Command::FFINDVAL},
    {"FREADHT", Command::FREADHT},
    {"FREADTH", Command::FREADTH},

    // --- LINKED (L) ---
    {"LPUSHHEAD", Command::LPUSHHEAD},
    {"LPUSH", Command::LPUSH},
    {"LPUSHBEFORE", Command::LPUSHBEFORE},
    {"LPUSHAFTER", Command::LPUSHAFTER},
    {"LPOPHEAD", Command::LPOPHEAD},
    {"LPOP", Command::LPOP},
    {"LPOPBEFORE", Command::LPOPBEFORE},
    {"LPOPAFTER", Command::LPOPAFTER},
    {"LDELVAL", Command::LDELVAL},
    {"LFINDVAL", Command::LFINDVAL},
    {"LREADHT", Command::LREADHT},
    {"LREADTH", Command::LREADTH},

    // --- STACK (S) ---
    {"SPUSH", Command::SPUSH},
    {"SPOP", Command::SPOP},
    {"SREAD", Command::SREAD},

    // --- QUEUE (Q) ---
    {"QPUSH", Command::QPUSH},
    {"QPOP", Command::QPOP},
    {"QREAD", Command::QREAD},

    // --- TREE (T) ---
    {"TINSERT", Command::TINSERT},
    {"TFIND", Command::TFIND},
    {"TDELETE", Command::TDELETE},
    {"TREAD", Command::TREAD},
    {"TCHECKAVL", Command::TCHECKAVL}
};


map<string,LinkedList> llists;
map<string,OneLinkedList> onellists;
map<string,Array> arrays;
map<string,Queue> queues;
map<string,Stack> stacks;
map<string,Tree> trees;

void write(string filename){
    ofstream fout;
    fout.open(filename);

    if (!fout.is_open()){
        throw runtime_error("Не удалось открыть файл при записи");
    }

    for (const auto& [name,obj] : arrays){
        fout << "Array ";
        fout << name << Array_read(obj) << endl;
    }
    for (const auto& [name,obj] : onellists){
        fout << "OneLinkedList ";
        fout << name << OneLinked_read(obj) << endl;
    }
    
    for (const auto& [datatype,obj] : llists){
        fout << "LinkedList ";
        fout << datatype <<Linked_read(obj) << endl;
    }
    for (const auto& [datatype,obj] : queues){
        fout << "Queue ";
        fout << datatype <<Queue_read(obj) << endl;
    }
    for (const auto& [datatype,obj] : stacks){
        fout << "Stack ";
        fout << datatype <<Stack_read(obj) << endl;
    }
    for (const auto& [datatype,obj] : trees){
        fout << "Tree ";
        fout << datatype <<Tree_read(obj) << endl;
    }
    fout << endl;

    fout.close();
}

void read(string filename){
    ifstream fin;
    fin.open(filename);

    if(!fin.is_open()){
        throw runtime_error("Не удалось открыть файл при чтении");
    }

    string datatype;
    while (fin >> datatype){
        string name;
        fin >> name;
        string line;
        getline(fin,line);
        stringstream ss(line);
        if (datatype == "Array"){
            Array arr = Array_create(10);
            string value;
            while (ss >> value){
                Array_push_back(arr,value);
            }
            arrays[name] = arr;
        }
        if (datatype == "OneLinkedList"){
            OneLinkedList onell = OneLinked_create();
            string value;
            while (ss >> value){
                OneLinked_push_back(onell,value);
            }
            onellists[name] = onell;
        }
        if (datatype == "LinkedList"){
            LinkedList ll = Linked_create();
            string value;
            while (ss >> value){
                Linked_push_back(ll,value);
            }
            llists[name] = ll;
        }
        if (datatype == "Queue"){
            Queue q = Queue_create();
            string value;
            while (ss >> value){
                Queue_push_back(q,value);
            }
            queues[name] = q;
        }
        if (datatype == "Stack"){
            Stack s = Stack_create();
            string value;
            while (ss >> value){
                Stack_push_back(s,value);
            }
            stacks[name] = s;
        }
        if (datatype == "Tree"){
            Tree t = Tree_create();
            string value;
            while (ss >> value){
                t.root = Tree_insert(t.root,stoi(value));
            }
            trees[name] = t;
        }
    }
    fin.close();
}

void make_change(const string& name, const vector<string>& data, Command command) {
    switch (command) {
        // --- ARRAY ---
        case Command::MCREATE: {
            Array arr = Array_create(10);
            arrays[name] = arr;
            break;
        }
        case Command::MPUSH: {
            for (auto& v : data)
                Array_push_back(arrays[name], v);
            break;
        }
        case Command::MPUSHIDX: {
            if (data.size() >= 2)
                Array_insert(arrays[name], (stoi(data[0])), data[1]);
            break;
        }
        case Command::MGET: {
            if (data.size() >= 1)
                cout << Array_get(arrays[name], stoi(data[0])) << endl;
            break;
        }
        case Command::MDELIDX: {
            if (data.size() >= 1)
                Array_erase(arrays[name], (stoi(data[0])));
            break;
        }
        case Command::MSETIDX: {
            if (data.size() >= 2)
                Array_set(arrays[name], (stoi(data[0])), data[1]);
            break;
        }
        case Command::MLENGTH: {
            cout << "Length: " << Array_length(arrays[name]) << endl;
            break;
        }
        case Command::MREAD: {
            cout << Array_read(arrays[name]) << endl;
            break;
        }

        // --- ONE-LINKED ---
        case Command::FPUSH: {
            for (auto& v : data)
                OneLinked_push_back(onellists[name], v);
            break;
        }
        case Command::FPUSHHEAD: {
            for (auto& v : data)
                OneLinked_push_front(onellists[name], v);
            break;
        }
        case Command::FPUSHBEFORE: {
            OneLinked_push_before(onellists[name], data[0],data[1]);
            break;
        }
        case Command::FPUSHAFTER: {
            OneLinked_push_after(onellists[name], data[0],data[1]);
            break;
        }
        case Command::FPOP: {
            OneLinked_pop_back(onellists[name]);
            break;
        }
        case Command::FPOPHEAD:{
            OneLinked_pop_front(onellists[name]);
            break;
        }
        case Command::FPOPAFTER: {
            OneLinked_pop_after(onellists[name], data[0]);
            break;
        }
        case Command::FPOPBEFORE: {
            OneLinked_pop_before(onellists[name],data[0]);
            break;
        }
        case Command::FDELVAL: {
            if (!data.empty())
                OneLinked_del_by_value(onellists[name], data[0]);
            break;
        }
        case Command::FFINDVAL: {
            cout << OneLinked_find_by_value(onellists[name],data[0]) << endl;
            break;
        }
        case Command::FREADHT: {
            cout << OneLinked_read(onellists[name]) << endl;
            break;
        }
        case Command::FREADTH: {
            cout << OneLinked_read_reverse(onellists[name]) << endl;
            break;
        }

        // --- LINKED (двусвязный) ---
        case Command::LPUSHHEAD: {
            for (auto& v : data)
                Linked_push_front(llists[name], v);
            break;
        }
        case Command::LPUSH: {
            for (auto& v : data)
                Linked_push_back(llists[name], v);
            break;
        }
        case Command::LPOPHEAD: {
            Linked_pop_front(llists[name]);
            break;
        }
        case Command::LPOP: {
            Linked_pop_back(llists[name]);
            break;
        }
        case Command::LPOPBEFORE:{
            Linked_pop_before(llists[name],data[0]);
            break;
        }
        case Command::LPOPAFTER:{
            Linked_pop_after(llists[name],data[0]);
            break;
        }
        case Command::LPUSHBEFORE: {
            Linked_push_before(llists[name],data[0],data[1]);
            break;
        }
        case Command::LPUSHAFTER: {
            Linked_push_after(llists[name],data[0],data[1]);
            break;
        }
        case Command::LDELVAL: {
            if (!data.empty())
                Linked_del_by_value(llists[name], data[0]);
            break;
        }
        case Command::LFINDVAL:{
            cout << Linked_find(llists[name],data[0]);
            break;
        }
        case Command::LREADHT: {
            cout << Linked_read(llists[name]) << endl;
            break;
        }
        case Command::LREADTH: {
            cout << Linked_read_reverse(llists[name]) << endl;
            break;
        }
        // --- STACK ---
        case Command::SPUSH: {
            for (auto& v : data)
                Stack_push_back(stacks[name], v);
            break;
        }
        case Command::SPOP: {
            Stack_pop(stacks[name]);
            break;
        }
        case Command::SREAD: {
            cout << Stack_read(stacks[name]) << endl;
            break;
        }

        // --- QUEUE ---
        case Command::QPUSH: {
            for (string v : data)
                Queue_push_back(queues[name], v);
            break;
        }
        case Command::QPOP: {
            Queue_pop(queues[name]);
            break;
        }
        case Command::QREAD: {
            cout << Queue_read(queues[name]) << endl;
            break;
        }

        // --- TREE ---
        case Command::TINSERT: {
            for (auto& v : data)
                trees[name].root = Tree_insert(trees[name].root, stoi(v));
            break;
        }
        case Command::TDELETE: {
            if (!data.empty())
                trees[name].root = Tree_delete(trees[name].root, stoi(data[0]));
            break;
        }
        case Command::TREAD: {
            cout << Tree_read(trees[name]) << endl;
            break;
        }

        // --- PRINT ---
        case Command::PRINT: {
            cout << "=== PRINT ===" << endl;

            if (arrays.count(name)) {
                cout << "Array " << name << ":" << Array_read(arrays[name]) << endl;
                break;
            }
            if (onellists.count(name)) {
                cout << "OneLinkedList " << name << ":" << OneLinked_read(onellists[name]) << endl;
                break;
            }
            if (llists.count(name)) {
                cout << "LinkedList " << name << " (Head→Tail):" 
                        << Linked_read(llists[name]) << endl;
                cout << "LinkedList " << name << " (Tail→Head):" 
                        << Linked_read_reverse(llists[name]) << endl;
                break;
            }
            if (queues.count(name)) {
                cout << "Queue " << name << ":" << Queue_read(queues[name]) << endl;
                break;
            }
            if (stacks.count(name)) {
                cout << "Stack " << name << ":" << Stack_read(stacks[name]) << endl;
                break;
            }
            if (trees.count(name)) {
                Tree_print(trees[name]);
                break;
            }

            cout << "Объект \"" << name << "\" не найден." << endl;
            break;
        }
        

        default:
            cout << "Неизвестная команда." << endl;
    }
}


void parsing(string filename,string query){
    read(filename);

    stringstream ss(query);

    string command;
    ss >> command;

    string name;
    ss >> name;
    vector<string> data;
    string temp;
    while (ss >> temp) {
        data.push_back(temp);
    }

    if (!command_map.count(command)){
        cout << "Неизвестная команда: " << command << endl;
        return;
    }

    make_change(name,data,command_map[command]);

    write(filename);
}

int main(int argc,char* argv[]) {
    string filename,query;

    if (argc != 5){
        return 0;
    }
    if (string(argv[1]) == "--file"){
        filename = argv[2];
    }

    if (string(argv[3]) == "--query"){
        query = argv[4];
    }

    try{
        parsing(filename,query);
    }
    catch(const exception& e){
        cout << e.what() << endl;
    }
}
