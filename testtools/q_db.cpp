#include <iostream>
#include <sqlite3.h>
#include <string>

void executeQuery(sqlite3* db, const std::string& query) {
    sqlite3_stmt* stmt;
    if (sqlite3_prepare_v2(db, query.c_str(), -1, &stmt, nullptr) == SQLITE_OK) {
        int colCount = sqlite3_column_count(stmt);

        while (sqlite3_step(stmt) == SQLITE_ROW) {
            for (int col = 0; col < colCount; ++col) {
                const char* text = (const char*)sqlite3_column_text(stmt, col);
                std::cout << (text ? text : "NULL") << " | ";
            }
            std::cout << std::endl;
        }
    } else {
        std::cerr << "Failed to execute query: " << sqlite3_errmsg(db) << std::endl;
    }
    sqlite3_finalize(stmt);
}

int main(int argc, char* argv[]) {
    if (argc != 2) {
        std::cerr << "Usage: " << argv[0] << " [-1 | -2 | -3 | -4 | -5 | -6]" << std::endl;
        return 1;
    }

    std::string param = argv[1];
    std::string query;
    std::string tableName;

    if (param == "-1") {
        query = "SELECT * FROM user_questionnaire_datas;";
        tableName = "user_questionnaire_datas";
    } else if (param == "-2") {
        query = "SELECT * FROM user_base_infos;";
        tableName = "user_base_infos";
    } else if (param == "-3") {
        query = "SELECT * FROM logins;";
        tableName = "logins";
    } else if (param == "-4") {
        query = "SELECT * FROM user_feedbacks;";
        tableName = "user_feedbacks";
    } else if (param == "-5") {
        query = "SELECT * FROM information_feedbacks;";
        tableName = "information_feedbacks";
    } else if (param == "-6") {
        query = "SELECT * FROM distribution_results;";
        tableName = "distribution_results";
    } else {
        std::cerr << "Invalid parameter. Use -1 for user_questionnaire_datas, -2 for user_base_infos, -3 for logins, -4 for user_feedbacks, -5 for information_feedbacks, and -6 for distribution_results." << std::endl;
        return 1;
    }

    sqlite3* db;
    int rc = sqlite3_open("../gorm.db", &db);

    if (rc) {
        std::cerr << "Can't open database: " << sqlite3_errmsg(db) << std::endl;
        return rc;
    } else {
        std::cout << "Opened database successfully, querying table: " << tableName << std::endl;
    }

    executeQuery(db, query);

    sqlite3_close(db);
    return 0;
}