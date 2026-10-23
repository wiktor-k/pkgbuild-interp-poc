# Maintainer: Torsten Keßler <tpkessler at archlinux dot org>
# Maintainer: Christian Heusel <gromit@archlinux.org>

pkgname=python-tensile
pkgver=7.2.4
pkgrel=2
pkgdesc="benchmark-driven backend libraries for general matrix-matrix multiplications"
arch=('any')
url='https://rocm.docs.amd.com/projects/Tensile/en/latest/'
license=('MIT')
depends=('python' 'python-msgpack' 'python-pyaml' 'python-joblib')
makedepends=('git' 'cmake' 'python-build' 'python-installer' 'python-wheel' 'python-setuptools')
_git='https://github.com/ROCm/rocm-libraries'
source=("rocm-libraries::git+$_git.git#tag=rocm-$pkgver")
sha256sums=('b476acbcd0f4017c800e4b05533e6dfb875bde32242729c8df557d4624379623')
_dirname="rocm-libraries/shared/tensile"

prepare() {
	cd ${_dirname}

	# Add support for gfx103X
	git cherry-pick -n 43a79e3aee2319377f2e69fe943b52a0c29215e8

	# Add support for gfx115X
	git cherry-pick -n 7bc1152aeaeaf9b6c0c7d3450be3d8afc571503b

	# https://src.fedoraproject.org/rpms/python-tensile/blob/rawhide/f/python-tensile.spec#_85
	sed -i -e 's@${Tensile_PREFIX}/bin/TensileGetPath@TensileGetPath@g' Tensile/cmake/TensileConfig.cmake
	sed -i -e 's@opt/rocm@usr@g' Tensile/Common.py
	sed -i -e 's@opt/rocm@usr@g' Tensile/Tests/yaml_only/test_config.py
	sed -i -e 's@globalParameters["IgnoreAsmCapCache"] = False@globalParameters["IgnoreAsmCapCache"] = True@' Tensile/Common.py
	sed -i -e 's@arguments["IgnoreAsmCapCache"] = args.IgnoreAsmCapCache@arguments["IgnoreAsmCapCache"] = True@' Tensile/TensileCreateLibrary.py
	sed -i -e 's@if not ignoreCacheCheck and derivedAsmCaps@if False and derivedAsmCaps@' Tensile/Common.py
}

build() {
	cd ${_dirname}
	python -m build --wheel --no-isolation
}

package() {
	cd ${_dirname}
	python -m installer --destdir="$pkgdir" dist/*.whl
	# cmake files are installed to the wrong location, /usr/cmake.
	# Move them to the right directory.
	mkdir -p "$pkgdir"/usr/lib/cmake/Tensile
	mv "$pkgdir"/usr/cmake/* "$pkgdir"/usr/lib/cmake/Tensile
	rm -rf "$pkgdir"/usr/cmake
	install -vDm644 LICENSE.md "$pkgdir"/usr/share/licenses/"$pkgname"/LICENSE
}
