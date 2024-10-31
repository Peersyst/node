// Import required libraries
const { ethers, parseUnits, Interface } = require('ethers');

// ABI for the mint function
const abi = [
    "function transfer(address to, uint256 amount) external",
    "function mint(address to, uint256 amount) external",
    "function approve(address spender, uint256 amount) external returns (bool)",
    "function balanceOf(address account) external view returns (uint256)",
    "function increaseAllowance(address spender, uint256 addedValue) external returns (bool)",
    "function allowance(address owner, address spender) external view returns (uint256)",
    "function addMinter(address minter) external"
];

async function mintTokens(
    contractAddress,
    recipientAddress,
    privateKey,
    rpcUrl = "http://127.0.0.1:8545"
) {
    try {
        // Connect to the network
        const provider = new ethers.JsonRpcProvider(rpcUrl);
        
        // Create a wallet instance
        const wallet = new ethers.Wallet(privateKey, provider);

        // Create contract instance
        const contract = new ethers.Contract(contractAddress, abi, wallet);


        // Initial balance
        const walletBalance = await contract.balanceOf(wallet.address);
        console.log("Wallet initial balance:", walletBalance.toString());

        const recipientBalance = await contract.balanceOf(recipientAddress);
        console.log("Recipient initial balance:", recipientBalance.toString());

        const contractBalance = await contract.balanceOf(contractAddress);
        console.log("Contract balance:", contractBalance.toString());

        const allowance = await contract.allowance(wallet.address, recipientAddress);
        console.log("Allowance:", allowance.toString());

        // Calculate the amount (0.01 * 10^18)
        const mintAmount = parseUnits("1000000000", 18);

        // await increaseAllowance(contract, recipientAddress, mintAmount);

        console.log('Minting tokens...');
        
        // Call mint function directly without approval
        const tx = await contract.mint(recipientAddress, mintAmount, {
           gasLimit: 4000000,
        });

        // Wait for transaction to be mined
        const receipt = await tx.wait();
        
        console.log('Minting successful!');
        console.log('Transaction hash:', receipt.transactionHash);

        // Check new balance
        const newRecipientBalance = await contract.balanceOf(recipientAddress);
        console.log("Recipient new balance:", newRecipientBalance.toString());
        
        return receipt;
    } catch (error) {
        console.error('Error minting tokens:', error);
        throw error;
    }
}

// Example usage:
async function main() {
    const CONTRACT_ADDRESS = "0xD4949664cD82660AaE99bEdc034a0deA8A0bd517";
    const RECIPIENT_ADDRESS = "0x9159C650e1D7E10a17c450eb3D50778aBA593D61";
    const PRIVATE_KEY = "7E1105DA629991DE37A309E0D51E7B39CDF661EC0DA2B748DA490CF6BF6BBD88";
    
    try {
        await mintTokens(
            CONTRACT_ADDRESS,
            RECIPIENT_ADDRESS,
            PRIVATE_KEY
        );
    } catch (error) {
        console.error('Main execution failed:', error);
    }
}

// Uncomment to run the script
main();