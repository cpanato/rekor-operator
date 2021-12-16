// Copyright 2021 The Sigstore Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controllers

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	operatorv1alpha1 "github.com/sigstore/rekor-operator/api/v1alpha1"
	"github.com/sigstore/rekor-operator/pkg/reconciler"
)

// RekorReconciler reconciles a Rekor object
type RekorReconciler struct {
	client.Client
	Log              logr.Logger
	ReconcileTimeout time.Duration
	Scheme           *runtime.Scheme
}

// +kubebuilder:rbac:groups=operator.rekor.dev,resources=rekors,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.rekor.dev,resources=rekors/status,verbs=get;update;patch

func (r *RekorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_, cancel := context.WithTimeout(ctx, reconciler.DefaultedLoopTimeout(r.ReconcileTimeout))
	defer cancel()
	_ = r.Log.WithValues("rekor", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

func (r *RekorReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		For(&operatorv1alpha1.Rekor{}).
		Complete(r)
}
