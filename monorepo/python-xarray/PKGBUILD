# Maintainer: Bruno Pagani <archange@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>

pkgname=python-xarray
pkgver=2026.09.0
pkgrel=1
pkgdesc="N-D labeled arrays and datasets in Python"
arch=(any)
url="https://xarray.pydata.org/"
license=(Apache-2.0)
depends=(
  python
  python-numpy
  python-packaging
  python-pandas
)
makedepends=(
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(
  python-aiobotocore
  python-bottleneck
  python-cftime
  python-dask
  python-distributed
  python-fsspec
  python-h5netcdf
  python-h5py
  python-hypothesis
  python-matplotlib
  python-netcdf4
  python-numexpr
  python-pint
  python-pyarrow
  python-pytest
  python-pytest-asyncio
  python-pytest-xdist
  python-scipy
  python-seaborn
)
optdepends=(
  'python-netcdf4: netCDF4 support'
  'python-h5netcdf: alternative for netCDF4 support'
  'python-scipy: interpolation features & fallback for netCDF3 support'
  # 'python-pydap: fallback for accessing OPeNDAP'
  # 'python-zarr: chunked, compressed N-dimensional arrays'
  'python-cftime: datetimes support for non-standard calendars or distant dates'
  'python-fsspec: remote filesystem support'
  'python-pooch: download and cache tutorial datasets'
  # 'python-pseudonetcdf: atmospheric science specific file formats support'
  # 'python-rasterio: GeoTiffs and other gridded raster datasets support'
  # 'python-iris: conversion to and from iris’ Cube objects'
  # 'python-cfgrib: mapping GRIB files to CDF4'
  'python-bottleneck: faster NaN-skipping and rolling window aggregations'
  'python-opt_einsum: faster tensor contractions'
  # 'python-numbagg: faster exponential rolling window operations'
  # 'python-flox: faster GroupBy reductions'
  'python-dask: parallel computation'
  'python-distributed: parallel computation'
  'python-matplotlib: plotting support'
  # 'python-cartopy: plotting of cartographic data'
  'python-seaborn: better color palettes for plots'
  # 'python-nc-time-axis: plotting of cftime.datetime objects'
  # 'python-sparse: sparse arrays support'
  'python-pint: units of measure support'
  'python-cupy: GPU array support'
  'python-hypothesis: strategies for testing code using xarray objects'
)
source=("https://github.com/pydata/xarray/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
b2sums=('ecc2ae0c17f9cfface94c5a94015ae9c92393afdeb35da0d2376e9e2a8b76e4970bc0f1bc8f6aa10c056c2863c1b20884c8d871a3981ad6b74195e0e5553375c')

build() {
  cd ${pkgname#python-}-$pkgver
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd ${pkgname#python-}-$pkgver
  local pytest_args=(
    --override-ini="addopts="
    -n=auto
    -W=ignore::pytest.PytestDeprecationWarning

    # pandas < 3 triggers this NumPy 2.5 warning when creating timedeltas.
    -W="default:The 'generic' unit for NumPy timedelta is deprecated:DeprecationWarning"

    # Segfaults with python-h5py installed
    --deselect=xarray/tests/test_backends.py::TestDask::test_save_mfdataset_compute_false_roundtrip

    # Fails with:
    # E   RuntimeError: NetCDF: Filter error: undefined filter encountered: (variable 'var2', group '/')
    --deselect=xarray/tests/test_backends.py::TestNetCDF4Data::test_compression_encoding
    --deselect=xarray/tests/test_backends.py::TestNetCDF4ViaDaskData::test_compression_encoding

    # Fails with:
    # E       Failed: DID NOT RAISE <class 'ValueError'>
    --deselect=xarray/tests/test_groupby.py::test_dask_da_groupby_quantile

    # Fails with:
    # pandas/_libs/arrays.pyx:103: NotImplementedError
    --deselect=xarray/tests/test_duck_array_ops.py::test_extension_array_attr

    # Fails with Python 3.14+ because inspect.signature() can now introspect
    # np.power, so no ValueError is raised
    --deselect=xarray/tests/test_dataarray.py::TestDataArray::test_curvefit_helpers
  )
  pytest "${pytest_args[@]}"
}

package() {
  cd ${pkgname#python-}-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  # Remove tests
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  rm -vr "$pkgdir/$site_packages/xarray/tests"
}
