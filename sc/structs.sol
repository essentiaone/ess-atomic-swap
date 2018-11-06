pragma solidity 0.4.25;

contract Structs {

    //Initiations struct for storing informations about unconfirmed order
    struct Initiations {
        address addressFrom;
        address addressTo;
        bool isShow;
        bool isRedeem;
        bool isInit;
        uint blockTimestamp;
        uint amount;
        bytes32 hashSecret;
    }

    //ConfirmedInitiations - struct for storing information about order that was already paid
    struct ConfirmedInitiations {
        address addressFrom;
        address addressTo;
        bool isShow;
        bool isRedeem;
        bool isInit;
        uint blockTimestamp;
        uint amount;
        bytes32 hashSecret;
    }
}
