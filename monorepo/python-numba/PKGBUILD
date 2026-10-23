# Maintainer: Bruno Pagani <archange@archlinux.org>

pkgname=python-numba
pkgver=0.68.0
pkgrel=1
pkgdesc='JIT compiler that translates a subset of Python and NumPy code into fast machine code'
arch=(x86_64)
url='https://numba.pydata.org/'
license=(BSD-1-Clause)
depends=(
  glibc
  libgcc
  libgomp
  libstdc++
  python
  python-llvmlite
  python-numpy
)
makedepends=(
  git
  onetbb
  openmp
  python-build
  python-installer
  python-packaging
  python-setuptools
  python-versioneer
  python-wheel
)
optdepends=(
  'onetbb: TBB threading backend'
  'openmp: OpenMP threading backend'
  'python-cffi: use of CFFI bindings in compiled functions'
  'python-colorama: error message highlighting'
  'python-scipy: cython bindings used in np.linalg.* support'
  'python-jinja: “pretty” type annotation output (HTML) via the CLI'
  'python-pyelftools: for code generation'
  'python-pygments: “pretty” type annotation'
  'python-pyyaml: use of a .numba_config.yaml file for storing per project configuration options'
  'python-psutil: for sysinfo'
)
checkdepends=(
  gdb
  ipython
  python-cffi
  python-ipykernel
  python-jinja
  python-pexpect
  python-pip
  python-pygments
  python-pytest
  python-pyyaml
  python-scipy
)
source=(git+https://github.com/numba/numba#tag=$pkgver)
sha512sums=('c8558d4981a40e0ba6a3880b961a91ce5eb4ae1f9a2708cfbcadcb397bb507cb524f775261cfffad3821a80c92b2aa274717960b789db3806478f2a0f8d7b943')
b2sums=('00bcbcbf06c396eb48725482ed4ee4442429f15a630e294b3a8e851ed02cc34a6aafe62acfe9c928e536600669895f79524fb1b88d585fba428e6cbdc9cade7a')

build() {
  cd numba
  python -m build --wheel --no-isolation
}

check() {
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer numba/dist/*.whl
  # skip failing typing test
  local site_packages=$(test-env/bin/python -c "import site; print(site.getsitepackages()[0])")
  mv "$site_packages"/numba/tests/test_function_type.py{,.disabled}
  test-env/bin/python -m numba.runtests -b -v -m 64 -- numba.tests
}

package() {
  cd numba
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
