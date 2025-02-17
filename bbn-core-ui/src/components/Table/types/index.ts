import type { ReactNode } from "react";

export type ColumnProps<T = unknown> = {
  key: string;
  header: string;
  headerClassName?: string;
  cellClassName?: string;
  render?: (value: unknown, row: T) => ReactNode;
  sorter?: (a: T, b: T) => number;
};

export type TableData = { id: string | number };

export type TableProps<T extends TableData> = ControlledTableProps & {
  data: T[];
  columns: ColumnProps<T>[];
  className?: string;
  wrapperClassName?: string;
  hasMore?: boolean;
  loading?: boolean;
  onLoadMore?: () => void;
  onRowSelect?: (row: T | null) => void;
  isRowSelectable?: (row: T) => boolean;
};

export interface ControlledTableProps {
  selectedRow?: string | number | null;
  defaultSelectedRow?: string | number | null;
  onSelectedRowChange?: (rowId: string | number | null) => void;
}
