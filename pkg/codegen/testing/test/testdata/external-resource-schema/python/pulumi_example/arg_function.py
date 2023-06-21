# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
import pulumi_random

__all__ = [
    'ArgFunctionResult',
    'AwaitableArgFunctionResult',
    'arg_function',
    'arg_function_output',
]

@pulumi.output_type
class ArgFunctionResult:
    def __init__(__self__, age=None):
        if age and not isinstance(age, int):
            raise TypeError("Expected argument 'age' to be a int")
        pulumi.set(__self__, "age", age)

    @property
    @pulumi.getter
    def age(self) -> Optional[int]:
        return pulumi.get(self, "age")


class AwaitableArgFunctionResult(ArgFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ArgFunctionResult(
            age=self.age)


def arg_function(name: Optional['pulumi_random.RandomPet'] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableArgFunctionResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('example::argFunction', __args__, opts=opts, typ=ArgFunctionResult).value

    return AwaitableArgFunctionResult(
        age=pulumi.get(__ret__, 'age'))


@_utilities.lift_output_func(arg_function)
def arg_function_output(name: Optional[pulumi.Input[Optional['pulumi_random.RandomPet']]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[ArgFunctionResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
