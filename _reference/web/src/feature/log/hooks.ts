import { useQuery } from '@tanstack/react-query'
import { logApi } from '@/api/log'
import { LogFilters } from '@/types/log'

// 获取日志数据
export const useLogs = (filters?: LogFilters) => {
    const query = useQuery({
        queryKey: ['logs', filters],
        queryFn: () => logApi.getLogData(filters),
        // 5分钟刷新一次数据
        refetchInterval: 5 * 60 * 1000,
        // 窗口重新获得焦点时刷新
        refetchOnWindowFocus: true,
        // 禁用重试，避免错误时过多请求
        retry: false,
    })

    return {
        ...query,
    }
}

// 获取日志详情
export const useLogDetail = (logId: number | null) => {
    const query = useQuery({
        queryKey: ['logDetail', logId],
        queryFn: () => {
            if (!logId) return null
            return logApi.getLogDetail(logId)
        },
        // 仅在有logId时启用查询
        enabled: !!logId,
        // 禁用自动重新获取
        refetchOnWindowFocus: false,
        refetchOnMount: false,
        refetchOnReconnect: false,
        refetchInterval: false,
        // 禁用重试
        retry: false,
    })

    return {
        ...query,
    }
} 