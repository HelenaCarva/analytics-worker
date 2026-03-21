// types.ts

export type AnalyticsApiRequest = {
  url: string;
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
  options?: {
    headers: Record<string, string>;
    body?: Record<string, unknown>;
  };
};

export type AnalyticsApiError = {
  code: number;
  message: string;
};

export type AnalyticsApiResponse<T> = {
  data: T;
  status: number;
  headers: Record<string, string>;
};

export type AnalyticsEvent = {
  type: 'PAGE_VIEW' | 'CLICK' | 'SUBMISSION';
  timestamp: Date;
  attributes: Record<string, string>;
};

export type AnalyticsCommand = {
  type: 'SEND_EVENT' | 'SEND_METRIC';
  event?: AnalyticsEvent;
  metric?: {
    name: string;
    value: number;
  };
};

export type AnalyticsMessage = {
  type: 'COMMAND' | 'HEARTBEAT';
  command?: AnalyticsCommand;
  heartbeat?: unknown;
};