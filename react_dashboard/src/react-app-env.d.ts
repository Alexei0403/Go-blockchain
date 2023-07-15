/// <reference types="react-scripts" />

type Transaction = {
  senderBlockchainAddress: string;
  recipientBlockchainAddress: string;
  value: number;
};

type Block = {
  timestamp: number;
  nonce: number;
  previousHash: string;
  transactions: Transaction[];
};

type WalletDetails = {
  blockchainAddress: string;
  privateKey: string;
  publicKey: string;
};

type WalletDetailsResponse = {
  blockchain_address: string;
  private_key: string;
  public_key: string;
};

type LocalError = {
  message: string;
} | null;

type Blockchain = {
  chain: Block[];
};
