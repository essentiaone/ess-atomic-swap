pragma solidity 0.4.25;

contract	AtomicSwap {

    struct Initiations {
        // uint ID; // need ID?
        address addressFrom;
        address addressTo;
        bool isShow;
        bool isRedeem;
        bool isInit;
    }

    struct ConfirmedInitiations {
        uint ID;
        address addressFrom;
        address addressTo;
        bool isShow;
        bool isRedeem;
    }

    mapping(address=>Initiations) public inits;

    mapping(address=>mapping(bytes20=>Initiations)) public confirmedInits;

    modifier isInitCreated(address _address) {
	    require(inits[_address].isInit == false);
	    _;
	}

    function addInit(address _addressFrom, address _addressTo) public isInitCreated(_addressFrom) returns(bool)
    {
        inits[_addressFrom].addressFrom = _addressFrom;
        inits[_addressFrom].addressTo = _addressTo;
        inits[_addressFrom].isShow = false;
        inits[_addressFrom].isRedeem = false;
        inits[_addressFrom].isInit = true;

        return true;
	}

	function getInit(address _addressOfInitiator) public view returns(address, address, bool, bool, bool) {
	    return (
	        inits[_addressOfInitiator].addressFrom,
	        inits[_addressOfInitiator].addressTo,
	        inits[_addressOfInitiator].isShow,
	        inits[_addressOfInitiator].isRedeem,
	        inits[_addressOfInitiator].isInit
	        );
	}
}
