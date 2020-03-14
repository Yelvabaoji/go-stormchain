pragma solidity =0.4.24 <0.6.3;
/**
 * @title ApplicationChain.sol
 * Stormchain can be registered on other blockchains as a
 * deligated application chain to support an application
 * using the tokens from the other chain.
 * This is the first step of building an application chain.
 */

contract AppChainBase {

    using SafeMath for uint256;

    address internal owner;
    mapping(address => uint) public admins;
    
    string public chainName;
    uint256 public balance = 0;
    uint256 public chainId;
    uint256 public period;
    uint256 public flushEpoch;
    uint256 public lastFlushedBlock = 0;

    string private genesisInfo;
    bool private genesisSet;

    mapping(uint256=>flushRound) public flushMapping;
    uint256[] public flushList;
    
    mapping(uint256=>address[]) public flushValidatorList;
    
    struct flushRound{
        uint256 flushId;
        address validator;
        uint256 blockNumber;
        string blockHash;
    }

    string extraData;
    uint256 tokenTotal;

    uint256 public FOUNDATION_MOAC_REQUIRED_AMOUNT = 5 * 10 ** 17;
    uint256 public FLUSH_AMOUNT = 5 * 10 ** 16;
    address public FOUNDATION_BLACK_HOLE_ADDRESS = 0x48328afC8dd45C1C252E7E883fc89bd17ddEe7c0;

    // construct a chain with no token
    constructor(string name, uint256 uniqueId, uint256 blockSec, uint256 flushNumber, address[] initial_validators, uint256 totalSupply) public payable {
        
        require(
            msg.value >= FOUNDATION_MOAC_REQUIRED_AMOUNT,
            "Not Enough MOAC to Create Application Chain"
        );

        owner = msg.sender;
        chainName = name;
        chainId = uniqueId;
        period = blockSec;
        flushEpoch = flushNumber;
        balance = msg.value;

        genesisSet = false;
        
        uint256 flushId = 0;
        flushMapping[flushId].flushId = flushId;
        flushMapping[flushId].validator = msg.sender;
        for (uint i=0; i<initial_validators.length; i++){
            flushValidatorList[flushId].push(initial_validators[i]);
        }
        flushMapping[flushId].blockNumber = 1;
        flushMapping[flushId].blockHash = "";      
        flushList.push(flushId);

        tokenTotal = totalSupply; 

        admins[msg.sender] = 1;
        
        // OPTIONAL: Give some gas fee to initial validators.
        // uint256 GAS_FEE = 2 * 10 ** 14;
        // uint256 gasNeededEach = balance / FLUSH_AMOUNT * GAS_FEE / initial_validators.length;
        // uint256 gasNeededTotal = gasNeededEach * initial_validators.length;
        // for (i=0; i<initial_validators.length; i++){
        //     initial_validators[i].transfer(gasNeededEach);
        // }
        // End of Option
    }

    function updateChainName(string name) public {
        require(admins[msg.sender] == 1, "Only Admins Can Update Chain Name.");
        chainName = name;
    }

    function updateFlushEpoch(uint256 newEpoch) public {
        require(admins[msg.sender] == 1, "Only Admins Can Update Flush Epoch.");
        require(newEpoch >= 360, "Flush Epoch Must be Equal to or Greater than 360.");
        flushEpoch = newEpoch;
    }

    function addAdmin(address admin) public {
        require(admins[msg.sender] == 1, "Only Admins Can Add Another Admin.");
        admins[admin] = 1;
    }

    function removeAdmin(address admin) public {
        require(admins[msg.sender] == 1, "Only Admins Can Add Another Admin.");
        require(admin != msg.sender, "Admins Cannot Remove Self.");
        admins[admin] = 0;
    }

    function addFund() public payable {
        balance += msg.value;
    }
    
    function withdrawFund(address recv, uint amount) public {
        require(admins[msg.sender] == 1);
        //require(admins[recv] == 1);
        require(amount <= balance);
        
        balance -= amount;
        recv.transfer(amount);
    }

    function flush(address[] current_validators, uint256 blockNumber, string blockHash) public {

        // do nothing if balance is less than 0.
        require(balance > 0, "Insufficient Amount");
        require(lastFlushedBlock + flushEpoch == blockNumber, "block number incorrect");

        uint256 flushId = flushList.length-1;
        for (uint i=0; i<flushValidatorList[flushId].length; i++){
            if ((flushValidatorList[flushId][i]==msg.sender ||
                admins[msg.sender] == 1 ) &&
                lastFlushedBlock + flushEpoch == blockNumber){

                flushId++;
                flushMapping[flushId].flushId = flushId;
                flushMapping[flushId].validator = msg.sender;
                for (uint j=0; j<current_validators.length; j++){
                        flushValidatorList[flushId].push(current_validators[j]);
                    }
                flushMapping[flushId].blockNumber = blockNumber;
                flushMapping[flushId].blockHash = blockHash;  
                flushList.push(flushId);
                
                // sent MOAC to FOUNDATION_BLACK_HOLE_ADDRESS;
                balance -= FLUSH_AMOUNT;
                FOUNDATION_BLACK_HOLE_ADDRESS.transfer(FLUSH_AMOUNT);
                lastFlushedBlock = flushMapping[flushId].blockNumber;
                return;
            }
        }
        require(1==1, "Only validators and admins can flush.");
    }

    function superflush(uint256 blockNumber, string blockHash) public {
        require(balance > 0, "Insufficient Amount");
        require(admins[msg.sender] == 1, "Only Admin Can Execute Superflush");
        uint256 flushId = flushList.length-1;
        flush(flushValidatorList[flushId], blockNumber, blockHash);
    }
    
    function setGenesisInfo(string genesis) public {
        require(admins[msg.sender] == 1, "Only Admins Can Set Genesis Info.");
        require(
            genesisSet == false,
            "Genesis Info Has Already Been Set."
        );
       genesisInfo = genesis;
        genesisSet = true;
    }

    function getGenesisInfo() public view returns (string) {
        return genesisInfo;
    }
}


/**
 * @title SafeMath
 * @dev Math operations with safety checks that revert on error
 */
library SafeMath {

  /**
  * @dev Multiplies two numbers, reverts on overflow.
  */
  function mul(uint256 _a, uint256 _b) internal pure returns (uint256) {
    // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
    // benefit is lost if 'b' is also tested.
    // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
    if (_a == 0) {
      return 0;
    }

    uint256 c = _a * _b;
    require(c / _a == _b);

    return c;
  }

  /**
  * @dev Integer division of two numbers truncating the quotient, reverts on division by zero.
  */
  function div(uint256 _a, uint256 _b) internal pure returns (uint256) {
    require(_b > 0); // Solidity only automatically asserts when dividing by 0
    uint256 c = _a / _b;
    // assert(_a == _b * c + _a % _b); // There is no case in which this doesn't hold

    return c;
  }

  /**
  * @dev Subtracts two numbers, reverts on overflow (i.e. if subtrahend is greater than minuend).
  */
  function sub(uint256 _a, uint256 _b) internal pure returns (uint256) {
    require(_b <= _a);
    uint256 c = _a - _b;

    return c;
  }

  /**
  * @dev Adds two numbers, reverts on overflow.
  */
  function add(uint256 _a, uint256 _b) internal pure returns (uint256) {
    uint256 c = _a + _b;
    require(c >= _a);

    return c;
  }

  /**
  * @dev Divides two numbers and returns the remainder (unsigned integer modulo),
  * reverts when dividing by zero.
  */
  function mod(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b != 0);
    return a % b;
  }
}
