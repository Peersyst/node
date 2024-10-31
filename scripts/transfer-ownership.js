const { ethers } = require("ethers");

const abi = [
    "function transfer(address to, uint256 amount) external",
    "function mint(address to, uint256 amount) external",
    "function approve(address spender, uint256 amount) external returns (bool)",
    "function balanceOf(address account) external view returns (uint256)",
    "function increaseAllowance(address spender, uint256 addedValue) external returns (bool)",
    "function allowance(address owner, address spender) external view returns (uint256)",
    "function addMinter(address minter) external",
    "function transferOwnership(address newOwner) external"
];

async function transferOwnership(contractAddress, privateKey, newOwner, rpcUrl = "http://127.0.0.1:8545") {

    try {
        // Connect to the network
        const provider = new ethers.JsonRpcProvider(rpcUrl);
        
        // Create a wallet instance
        const wallet = new ethers.Wallet(privateKey, provider);

        // Create contract instance
        const contract = new ethers.Contract(contractAddress, abi, wallet);

        console.log('Transferring ownership...');
        
        // Call mint function directly without approval
        const tx = await contract.transferOwnership(newOwner, {
           gasLimit: 4000000,
        });

        // Wait for transaction to be mined
        const receipt = await tx.wait();

        console.log('Transfer ownership successful!');
        console.log('Transaction hash:', receipt.transactionHash);
        
        return receipt;
    } catch (error) {
        console.error('Error minting tokens:', error);
        throw error;
    }
}

async function main() {
    const CONTRACT_ADDRESS = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517";
    const PRIVATE_KEY = "7E1105DA629991DE37A309E0D51E7B39CDF661EC0DA2B748DA490CF6BF6BBD88";
    const NEW_OWNER = "0x9159C650e1D7E10a17c450eb3D50778aBA593D61";
    await transferOwnership(CONTRACT_ADDRESS, PRIVATE_KEY, NEW_OWNER);
}

main();