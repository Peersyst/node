// Import required libraries
const { ethers, parseUnits } = require('ethers');

// ABI for the mint function
const abi = [
    "function transfer(address to, uint256 amount) external",
    "function mint(address to, uint256 amount) external",
    "function burn(uint256 amount) external",
    "function approve(address spender, uint256 amount) external returns (bool)",
    "function balanceOf(address account) external view returns (uint256)",
    "function increaseAllowance(address spender, uint256 addedValue) external returns (bool)",
    "function allowance(address owner, address spender) external view returns (uint256)",
    "function owner() external view returns (address)"
];

async function getOwner(contractAddress, privateKey, rpcUrl = "http://127.0.0.1:8545") {
     try {
        // Connect to the network
        const provider = new ethers.JsonRpcProvider(rpcUrl);
        
        // Create a wallet instance
        const wallet = new ethers.Wallet(privateKey, provider);

		// console.log("Wallet address: ", wallet.address);
        
        // Create contract instance
        const contract = new ethers.Contract(contractAddress, abi, wallet);

        console.log('Querying owner...');
        
        // Call mint function directly without approval
        const owner = await contract.owner();
        console.log("Owner:", owner);
        
        return owner;
    } catch (error) {
        console.error('Error querying owner:', error);
        throw error;
    }
}

async function main() {
    const CONTRACT_ADDRESS = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517";
    const PRIVATE_KEY = "7E1105DA629991DE37A309E0D51E7B39CDF661EC0DA2B748DA490CF6BF6BBD88";
    await getOwner(CONTRACT_ADDRESS, PRIVATE_KEY);
}

main();