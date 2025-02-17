import { networks } from "bitcoinjs-lib";

import { Network } from "@/core/types";

export function validateAddress(network: Network, address: string): void {
  if (network === Network.MAINNET && !address.startsWith("bc1")) {
    throw new Error("Incorrect address prefix for Mainnet. Expected address to start with 'bc1'.");
  } else if ([Network.SIGNET, Network.TESTNET].includes(network) && !address.startsWith("tb1")) {
    throw new Error("Incorrect address prefix for Testnet / Signet. Expected address to start with 'tb1'.");
  } else if (![Network.MAINNET, Network.SIGNET, Network.TESTNET].includes(network)) {
    throw new Error(`Unsupported network: ${network}. Please provide a valid network.`);
  }
}

export const toNetwork = (network: Network): networks.Network => {
  switch (network) {
    case Network.MAINNET:
      return networks.bitcoin;
    case Network.TESTNET:
    case Network.SIGNET:
      return networks.testnet;
    default:
      throw new Error("Unsupported network");
  }
};
