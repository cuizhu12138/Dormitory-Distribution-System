#include <iostream>
#include <sqlite3.h>
#include <string>

void executeDeleteTable(sqlite3* db, const std::string& tableName) {
    std::string query = "DROP TABLE IF EXISTS " + tableName + ";";
    char* errMsg = nullptr;
    
    int rc = sqlite3_exec(db, query.c_str(), nullptr, nullptr, &errMsg);
    
    if (rc != SQLITE_OK) {
        std::cerr << "Failed to delete table " << tableName << ": " << errMsg << std::endl;
        sqlite3_free(errMsg);
    } else {
        std::cout << "Deleted table successfully: " << tableName << std::endl;
    }
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " [-1 | -2 | -3 | -4 | -5 | -6]" << std::endl;
        return 1;
    }

    std::string param = argv[1];
    std::string tableName;

    if (param == "-1") {
        tableName = "user_questionnaire_datas";
    } else if (param == "-2") {
        tableName = "user_base_infos";
    } else if (param == "-3") {
        tableName = "logins";
    } else if (param == "-4") {
        tableName = "user_feedbacks";
    } else if (param == "-5") {
        tableName = "information_feedbacks";
    } else if (param == "-6") {
        tableName = "distribution_results";
    } else {
        std::cerr << "Invalid parameter. Use -1 for user_questionnaire_datas, -2 for user_base_infos, -3 for logins, -4 for user_feedbacks, -5 for information_feedbacks, and -6 for distribution_results." << std::endl;
        return 1;
    }

    std::string dbPath = "../gorm.db"; // 上级目录下的gorm.db路径

    sqlite3* db;
    int rc = sqlite3_open(dbPath.c_str(), &db);

    if (rc) {
        std::cerr << "Can't open database: " << sqlite3_errmsg(db) << std::endl;
        return rc;
    } else {
        std::cout << "Opened database successfully, attempting to delete table: " << tableName << std::endl;
    }

    executeDeleteTable(db, tableName);

    sqlite3_close(db);
    return 0;
}