import { createNanoEvents } from "nanoevents";

import { Wallet } from "@/core/Wallet";
import type { IConnector, IProvider } from "@/core/types";

export interface ConnectorEvents<P extends IProvider> {
  connecting: (message?: string) => void;
  connect: (wallet: Wallet<P>) => void;
  disconnect: (wallet: Wallet<P>) => void;
  error: (error: Error) => void;
}

export class WalletConnector<N extends string, P extends IProvider, C> implements IConnector<N, P, C> {
  private _connectedWallet: Wallet<P> | null = null;
  private _ee = createNanoEvents<ConnectorEvents<P>>();

  constructor(
    public readonly id: N,
    public readonly name: string,
    public readonly icon: string,
    public readonly wallets: Wallet<P>[],
    public readonly config: C,
  ) {}

  get connectedWallet() {
    return this._connectedWallet;
  }

  async connect(wallet: string | Wallet<P>) {
    try {
      const selectedWallet = typeof wallet === "string" ? this.wallets.find((w) => w.id === wallet) : wallet;

      if (!selectedWallet) {
        throw new Error("Wallet not found");
      }
      this._ee.emit("connecting", `Connecting ${selectedWallet.name}`);

      await selectedWallet.connect();
      this._connectedWallet = selectedWallet;
      this._ee.emit("connect", this._connectedWallet);

      return this.connectedWallet;
    } catch (e: any) {
      this._ee.emit("error", e);
      return null;
    }
  }

  async disconnect() {
    if (this._connectedWallet) {
      this._ee.emit("disconnect", this._connectedWallet);
      this._connectedWallet = null;
    }
  }

  clone() {
    return new WalletConnector(this.id, this.name, this.icon, this.wallets, this.config);
  }

  on<K extends keyof ConnectorEvents<P>>(name: K, handler: ConnectorEvents<P>[K]) {
    return this._ee.on(name, handler);
  }
}
