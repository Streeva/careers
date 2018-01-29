
namespace streeva.careers.coreengineer.reactive
{
    using System;

    public interface IDataSource
    {
        IObservable<Data> DataFeed { get; }
    }
}