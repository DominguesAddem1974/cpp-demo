#include <gtest/gtest.h>

#include "function.h"

TEST(testCase, test1) {
    EXPECT_EQ(Add(2, 3), 5);
}

int main(int argc, char **argv) {
    testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}