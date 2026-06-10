// src/feature/model/hooks.ts
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { modelApi } from "@/api/model";
import { useState } from "react";
import { ModelCreateRequest } from "@/types/model";
import { toast } from "sonner";
import { ApiError } from "@/api/index";

// Get all models
export const useModels = () => {
  return useQuery({
    queryKey: ["models"],
    queryFn: () => modelApi.getModels(),
    staleTime: 1000 * 60 * 5, // 5 minutes
  });
};

// Get model sets data
export const useModelSets = () => {
  return useQuery({
    queryKey: ["modelSets"],
    queryFn: () => modelApi.getModelSets(),
    staleTime: 1000 * 60 * 5, // 5 minutes
  });
};

// Get builtin model configs grouped by channel type
export const useChannelBuiltinModels = () => {
  return useQuery({
    queryKey: ["channelBuiltinModels"],
    queryFn: () => modelApi.getChannelBuiltinModels(),
    staleTime: 1000 * 60 * 30,
  });
};

// Get a specific model
export const useModel = (model: string) => {
  return useQuery({
    queryKey: ["model", model],
    queryFn: () => modelApi.getModel(model),
    enabled: !!model,
    staleTime: 1000 * 60 * 5, // 5 minutes
  });
};

// Create a new model
export const useCreateModel = () => {
  const queryClient = useQueryClient();
  const [error, setError] = useState<ApiError | null>(null);

  const mutation = useMutation({
    mutationFn: (data: ModelCreateRequest) => {
      return modelApi.createModel(data.model, data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["models"] });
      setError(null);
    },
    onError: (err: ApiError) => {
      setError(err);
      toast.error(err.message);
    },
  });

  return {
    createModel: mutation.mutate,
    isLoading: mutation.isPending,
    error,
    clearError: () => setError(null),
  };
};

// Update an existing model
export const useUpdateModel = () => {
  const queryClient = useQueryClient();
  const [error, setError] = useState<ApiError | null>(null);

  const mutation = useMutation({
    mutationFn: ({
      model,
      data,
    }: {
      model: string;
      data: Omit<ModelCreateRequest, "model">;
    }) => {
      return modelApi.updateModel(model, data);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["models"] });
      setError(null);
    },
    onError: (err: ApiError) => {
      setError(err);
      toast.error(err.message);
    },
  });

  return {
    updateModel: mutation.mutate,
    isLoading: mutation.isPending,
    error,
    clearError: () => setError(null),
  };
};

// Delete a model
export const useDeleteModel = () => {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (model: string) => {
      return modelApi.deleteModel(model);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["models"] });
      toast.success("Model deleted successfully");
    },
    onError: (err: ApiError) => {
      toast.error(err.message);
    },
  });

  return {
    deleteModel: mutation.mutate,
    isLoading: mutation.isPending,
  };
};
